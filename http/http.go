package http

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// PostRequest sends a POST method request
// and returns the response.
func PostRequest(url string, headers map[string]string, body string) (string, error) {
	// TODO: Add POST timeout here.
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	for headerKey, headerValue := range headers {
		req.Header.Set(headerKey, headerValue)
	}
	res, err := client.Do(req)
	if res != nil {
		defer func() { _ = res.Body.Close() }()
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
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if len(opt) != 0 {
		option.copyWith(opt[0])
	}
	if option.Cookie != "" {
		req.Header.Add("cookie", option.Cookie)
	}
	for key, value := range option.Header {
		req.Header.Set(key, value)
	}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
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
