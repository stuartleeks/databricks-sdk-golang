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

// DataBrickClient is the client to init for communicating with DataBricks, Python source:
// https://github.com/databricks/databricks-cli/blob/master/databricks_cli/sdk/api_client.py
// based on commit: v0.8.6 (cb22ab21327b7d1c85bb45404f7b1115e1fa9dcf)
type DataBrickClient struct {
	User               string
	Password           string
	Host               string
	Token              string
	DefaultHeaders     map[string]string
	InsecureSkipVerify bool
	TimeoutSeconds     int
}

func (dbc *DataBrickClient) performQuery(
	method, path string,
	data map[string]interface{},
	headers map[string]string,
) ([]byte, error) {

	/*
		parsed_url = urlparse(host)
		scheme = parsed_url.scheme
		hostname = parsed_url.hostname
		self.url = "%s://%s/api/%s" % (scheme, hostname, apiVersion)
	*/
	parsedURL, err := url.Parse(dbc.Host)
	requestURL := fmt.Sprintf("%s://%s/api/%s%s", parsedURL.Scheme, parsedURL.Host, apiVersion, path)

	/*
		if user is not None and password is not None:
				encoded_auth = (user + ":" + password).encode()
				user_header_data = "Basic " + base64.standard_b64encode(encoded_auth).decode()
				auth = {'Authorization': user_header_data, 'Content-Type': 'text/json'}
		elif token is not None:
				auth = {'Authorization': 'Bearer {}'.format(token), 'Content-Type': 'text/json'}
		else:
				auth = {}
	*/
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

	/*
		user_agent = {'user-agent': 'databricks-cli-{v}-{c}'.format(v=databricks_cli_version,
																																c=command_name)}
		self.default_headers = {}
		self.default_headers.update(auth)
		self.default_headers.update(default_headers)
		self.default_headers.update(user_agent)
	*/
	userAgent := map[string]string{
		"user-agent": fmt.Sprintf("databricks-sdk-golang-%s", sdkVersion),
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

	/*
		if headers is None:
				headers = self.default_headers
		else:
				tmp_headers = copy.deepcopy(self.default_headers)
				tmp_headers.update(headers)
				headers = tmp_headers
	*/
	if headers == nil || len(headers) == 0 {
		// Do Nothing
	} else {
		for k, v := range headers {
			requestHeaders[k] = v
		}
	}

	/*
		with warnings.catch_warnings():
				warnings.simplefilter("ignore", exceptions.InsecureRequestWarning)
				if method == 'GET':
						translated_data = {k: _translate_boolean_to_query_param(data[k]) for k in data}
						resp = self.session.request(method, self.url + path, params = translated_data,
								verify = self.verify, headers = headers)
				else:
						resp = self.session.request(method, self.url + path, data = json.dumps(data),
								verify = self.verify, headers = headers)
	*/
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

	/*
		try:
				resp.raise_for_status()
		except requests.exceptions.HTTPError as e:
				message = e.args[0]
				try:
						reason = pprint.pformat(json.loads(resp.text), indent=2)
						message += '\n Response from server: \n {}'.format(reason)
				except ValueError:
						pass
				raise requests.exceptions.HTTPError(message, response=e.response)
		return resp.json()
	*/
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
	/*
		assert not isinstance(value, list), 'GET parameters cannot pass list of objects'
		if isinstance(value, bool):
				if value:
						return 'true'
				else:
						return 'false'
		return value
	*/
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
