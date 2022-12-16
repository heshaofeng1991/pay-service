package tool

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func SendJSONRequestToAPI(url string, method string, postData interface{},
	headers map[string]string,
) ([]byte, error) {
	var (
		req           *http.Request
		jsonStr, data []byte
		err           error
	)

	// 30秒超时
	transport := &http.Transport{
		TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{},
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
	}

	var body io.Reader

	// 构造request
	if postData != nil {
		r, ok := postData.(string)
		if ok {
			body = strings.NewReader(r)
		} else {
			jsonStr, err = json.Marshal(postData)
			if err != nil {
				return data, err
			}

			body = strings.NewReader(string(jsonStr))
		}
	}

	req, err = http.NewRequestWithContext(context.Background(), method, url, body)
	if err != nil {
		return data, err
	}

	// 设置Header
	ProcessHTTPHeader(req, headers)

	req.Close = true

	// 发起请求
	rsp, err := client.Do(req)
	if err != nil {
		return data, err
	}

	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return data, errors.New("http 请求失败")
	}

	// 解析返回.
	data, err = ioutil.ReadAll(rsp.Body)
	if err == nil {
		return data, nil
	}

	if strings.Contains(err.Error(), "unexpected EOF") {
		return data, nil
	}

	return nil, err
}

func SendGetRequest(url string, params map[string]interface{},
	headers map[string]string,
) ([]byte, error) {
	var (
		req  *http.Request
		data []byte
		err  error
	)

	client := &http.Client{}

	req, err = http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return data, err
	}

	// 设置Header.
	ProcessHTTPHeader(req, headers)

	// 设置Param.
	ProcessHTTPParam(req, params)

	rsp, err := client.Do(req)
	if err != nil {
		return data, err
	}

	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return data, errors.New("http 请求失败")
	}

	return io.ReadAll(rsp.Body)
}

func SendRequestToBaseAPI(url, method string, body io.Reader, headers map[string]string) ([]byte, error) {
	var (
		req  *http.Request
		data []byte
		err  error
		resp *http.Response
	)

	// 30秒超时
	// transport := &http.Transport{
	// 	TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{},
	// }

	client := &http.Client{
		Timeout: 30 * time.Second,
		// Transport: transport,
	}

	req, err = http.NewRequestWithContext(context.Background(), method, url, body)
	if err != nil {
		return data, err
	}

	// set Header
	ProcessHTTPHeader(req, headers)

	resp, err = client.Do(req)
	if err != nil {
		return data, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return data, errors.New("http 请求失败")
	}

	return io.ReadAll(resp.Body)
}

func ProcessHTTPHeader(req *http.Request, headers map[string]string) {
	if len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
}

func ProcessHTTPParam(req *http.Request, params map[string]interface{}) {
	query := req.URL.Query()

	if len(params) > 0 {
		for key, value := range params {
			switch valueType := value.(type) {
			case string:
				query.Add(key, valueType)
			case int:
				val := strconv.Itoa(valueType)
				query.Add(key, val)
			case []string:
				for _, item := range valueType {
					query.Add(key, item)
				}
			}
		}
	}

	req.URL.RawQuery = query.Encode()
}
