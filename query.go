package databricks

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func (dbc *DataBrickClient) performQuery(method, path string, data map[string]interface{}, headers map[string]string) ([]byte, error) {

	parsedURL, err := url.Parse(dbc.Host)
	requestURL := fmt.Sprintf("%s://%s/api/%s%s", parsedURL.Scheme, parsedURL.Host, APIVersion, path)

	auth := make(map[string]string)
	if dbc.User != "" && dbc.Password != "" {
		encodedAuth := []byte(dbc.User + ":" + dbc.Password)
		userHeaderData := "Basic " + base64.StdEncoding.EncodeToString(encodedAuth)
		auth["Authorization"] = userHeaderData
		auth["Content-Type"] = "text/json"
	} else if dbc.Token != "" {
		auth["Authorization"] = "Bearer " + dbc.Token
		auth["Content-Type"] = "text/json"
	} else {
		// Do Nothing
	}

	userAgent := map[string]string{
		"user-agent": fmt.Sprintf("databricks-sdk-golang-%s", SdkVersion),
	}

	requestHeaders := make(map[string]string)

	for k, v := range auth {
		requestHeaders[k] = v
	}
	for k, v := range dbc.DefaultHeaders {
		requestHeaders[k] = v
	}
	for k, v := range userAgent {
		requestHeaders[k] = v
	}

	if headers == nil || len(headers) == 0 {
		// Do Nothing
	} else {
		for k, v := range headers {
			requestHeaders[k] = v
		}
	}

	var requestBody []byte
	if method == "GET" {
		params := url.Values{}
		for k, v := range data {
			translatedValue, err := dbc.translateBooleanToQueryParam(v)
			if err != nil {
				return nil, err
			}
			params.Add(k, translatedValue)
		}
		requestURL += "?" + params.Encode()
	} else {
		bodyBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		requestBody = bodyBytes
	}

	if dbc.TimeoutSeconds == 0 {
		dbc.TimeoutSeconds = 10
	}
	client := http.Client{
		Timeout: time.Duration(time.Duration(dbc.TimeoutSeconds) * time.Second),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: dbc.InsecureSkipVerify,
			},
		},
	}

	request, err := http.NewRequest(method, requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("Response from server (%d) %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func (dbc *DataBrickClient) translateBooleanToQueryParam(v interface{}) (string, error) {
	switch t := v.(type) {
	case []interface{}:
		return "", fmt.Errorf("GET parameters cannot pass list of objects")
	case map[string]interface{}:
		return "", fmt.Errorf("GET parameters cannot pass a map")
	case bool:
		return strconv.FormatBool(t), nil
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64), nil
	case int64:
		return strconv.FormatInt(t, 10), nil
	case string:
		return t, nil
	default:
		return "", fmt.Errorf("GET parameters unknown object")
	}
}
