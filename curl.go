package curl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

// request method
type MethodType string

type ContentType string

const (
	// method
	Get    MethodType = "GET"
	Post   MethodType = "POST"
	Put    MethodType = "PUT"
	Patch  MethodType = "PATCH"
	Delete MethodType = "DELETE"

	// content_type
	JsonType  ContentType = "application/json"
	FormType  ContentType = "application/x-www-form-urlencoded"
	OtherType ContentType = "null" // need set header and body

	// timeout const
	timeout = 30 * time.Second // second
)

// Client struct
type Client struct {
	host    string
	headers map[string]string
	cookies []*http.Cookie
	timeout time.Duration

	// 'OtherType', you need set value
	body string
}

// NewClient new client instance
func NewClient(host string) *Client {
	return &Client{
		host:    host,
		headers: map[string]string{},
		timeout: timeout,
	}
}

// SetHeaders set request headers
func (c *Client) SetHeaders(headers map[string]string) *Client {
	c.headers = headers
	return c
}

// SetCookies set request cookies
func (c *Client) SetCookies(cookies []*http.Cookie) *Client {
	c.cookies = cookies
	return c
}

// SetTimeout set request timeout
func (c *Client) SetTimeout(timeout time.Duration) *Client {
	c.timeout = timeout
	return c
}

// SetBody set request body
func (c *Client) SetBody(body string) *Client {
	c.body = body
	return c
}

// Get request get, content_type for "application/json"
func (c *Client) Get(path string, data interface{}) ([]byte, error) {
	return c.Curl(path, Get, data, JsonType)
}

// Delete request delete, content_type for "application/json"
func (c *Client) Delete(path string, data interface{}) ([]byte, error) {
	return c.Curl(path, Delete, data, JsonType)
}

// Post request post, content_type for "application/json"
func (c *Client) Post(path string, data interface{}) ([]byte, error) {
	return c.Curl(path, Post, data, JsonType)
}

// PostByForm request post, content_type for "application/x-www-form-urlencoded"
func (c *Client) PostByForm(path string, data interface{}) ([]byte, error) {
	return c.Curl(path, Post, data, FormType)
}

// Put request put, content_type for "application/json"
func (c *Client) Put(path string, data interface{}) ([]byte, error) {
	return c.Curl(path, Put, data, JsonType)
}

// PutByForm request put, content_type for "application/x-www-form-urlencoded"
func (c *Client) PutByForm(path string, data interface{}) ([]byte, error) {
	return c.Curl(path, Put, data, FormType)
}

// Patch request patch, content_type for "application/json"
func (c *Client) Patch(path string, data interface{}) ([]byte, error) {
	return c.Curl(path, Patch, data, JsonType)
}

// PatchByForm request patch, content_type for "application/x-www-form-urlencoded"
func (c *Client) PatchByForm(path string, data interface{}) ([]byte, error) {
	return c.Curl(path, Patch, data, JsonType)
}

// Curl url request
func (c *Client) Curl(path string, method MethodType, data interface{}, cType ContentType) ([]byte, error) {
	// request client, parse url and body
	client := &http.Client{Timeout: 30 * time.Second}
	reqUrl, body, err := c.parse(path, method, data, cType)
	if err != nil {
		return nil, err
	}

	// request data
	switch cType {
	case JsonType:
		c.headers["Content-Type"] = "application/json"
	case FormType:
		c.headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	// request
	req, err := http.NewRequest(string(method), reqUrl, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	// header option
	for hName, header := range c.headers {
		req.Header.Add(hName, header)
	}

	// cookie option
	for _, cookie := range c.cookies {
		req.AddCookie(cookie)
	}

	// do request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyByte, nil
}

// parse url and data
func (c *Client) parse(path string, method MethodType, data interface{}, cType ContentType) (string, string, error) {
	var body string
	reqUrl := c.getApiUrl(path)

	// request data
	if data == nil {
		return reqUrl, body, nil
	}

	rv := reflect.ValueOf(data)
	switch rv.Kind() {
	case reflect.String:
		return reqUrl, fmt.Sprintf("%v", data), nil
	case reflect.Map:
		switch method {
		case Get, Delete:
			params, err := getUrlValues(data)
			if err != nil {
				return "", "", err
			}

			reqUrlObj, err := url.Parse(reqUrl)
			if err != nil {
				return "", "", err
			}
			reqUrlObj.RawQuery = params.Encode()
			reqUrl = reqUrlObj.String()
		case Post, Put, Patch:
			switch cType {
			case JsonType:
				dataBytes, err := json.Marshal(data)
				if err != nil {
					return "", "", err
				}
				body = string(dataBytes)
			case FormType:
				params, err := getUrlValues(data)
				if err != nil {
					return "", "", err
				}
				body = params.Encode()
			default:
				body = c.body
			}
		}
	}

	return reqUrl, body, nil
}

// get api_url
func (c *Client) getApiUrl(path string) string {
	return fmt.Sprintf("%s%s", c.host, path)
}

// 设置 url values
func getUrlValues(data interface{}) (*url.Values, error) {
	dataMap, dataOk := data.(map[string]string)
	if data != nil && dataOk == false {
		return nil, errors.New("the request type should use 'map[string]string'")
	}

	params := make(url.Values)
	for key, value := range dataMap {
		params.Set(key, value)
	}

	return &params, nil
}
