package http

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// PostRequest sends a POST method request
// and returns the response.
func PostRequest(url string, opt ...Option) (string, error) {
	// TODO: Add POST timeout here.
	client := &http.Client{}
	option := defaultOption()
	if len(opt) != 0 {
		option.copyWith(opt[0])
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(option.Body))
	if err != nil {
		return "", err
	}
	if option.Cookie != "" {
		req.Header.Add("cookie", option.Cookie)
	}
	for headerKey, headerValue := range option.Header {
		req.Header.Set(headerKey, headerValue)
	}

	query := req.URL.Query()
	for queryKey, queryValue := range option.Query {
		query.Add(queryKey, queryValue)
	}
	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	if res != nil {
		defer func() { _ = res.Body.Close() }()
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%d:%s", res.StatusCode, res.Status)
	}
	if err != nil {
		return "", err
	}
	result, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// GetRequest sends a GET method request
// and returns the response.
func GetRequest(url string, opt ...Option) ([]byte, error) {
	client := &http.Client{}
	var resp *http.Response
	var ret []byte
	option := defaultOption()
	if len(opt) != 0 {
		option.copyWith(opt[0])
	}
	req, err := http.NewRequest("GET", url, strings.NewReader(option.Body))
	if err != nil {
		return nil, err
	}
	if option.Cookie != "" {
		req.Header.Add("cookie", option.Cookie)
	}
	for key, value := range option.Header {
		req.Header.Set(key, value)
	}

	query := req.URL.Query()
	for queryKey, queryValue := range option.Query {
		query.Add(queryKey, queryValue)
	}
	req.URL.RawQuery = query.Encode()

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil fmt.Errorf("%d:%s", resp.StatusCode, resp.Status)
	}
	defer func() { _ = resp.Body.Close() }()
	ret, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err == nil && resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%s[%d]", resp.Status, resp.StatusCode)
	}
	return ret, err
}
