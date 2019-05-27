package databricks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func performQuery(
	option DBClientOption,
	method, path string,
	data map[string]interface{}, headers map[string]string) ([]byte, error) {

	requestURL, err := option.getRequestURI(path)
	if err != nil {
		return nil, err
	}

	requestHeaders := option.getDefaultHeaders()

	if headers != nil && len(headers) > 0 {
		for k, v := range headers {
			requestHeaders[k] = v
		}
	}

	var requestBody []byte
	if method == "GET" {
		params := url.Values{}
		for k, v := range data {
			translatedValue, err := translateBooleanToQueryParam(v)
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

	client := option.getHTTPClient()

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

func translateBooleanToQueryParam(v interface{}) (string, error) {
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
