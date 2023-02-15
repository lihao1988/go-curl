package curl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
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
	OtherType ContentType = "you set data"

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

// CurlForJson url request for json
func (c *Client) CurlForJson(uri string, method MethodType, data map[string]interface{}) ([]byte, error) {
	dataJsonByte, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return c.Curl(uri, method, string(dataJsonByte), JsonType)
}

// Curl url request
func (c *Client) Curl(uri string, method MethodType, data interface{}, cType ContentType) ([]byte, error) {
	// request client
	client := &http.Client{Timeout: 30 * time.Second}

	// request data
	var body io.Reader
	switch cType {
	case JsonType:
		c.headers["Content-Type"] = "application/json"
		body = strings.NewReader(fmt.Sprintf("%v", data))
	case FormType:
		formDataMap, dataOk := data.(map[string]string)
		if data != nil && dataOk {
			return nil, errors.New("the request type should use [map]")
		}
		formData := make(url.Values)
		c.headers["Content-Type"] = "application/x-www-form-urlencoded"
		for key, value := range formDataMap {
			formData.Set(key, value)
		}
		body = strings.NewReader(formData.Encode())
	default:
		body = strings.NewReader(c.body)
	}

	// request
	req, err := http.NewRequest(string(method), c.getApiUrl(uri), body)
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

// get api_url
func (c *Client) getApiUrl(uri string) string {
	return fmt.Sprintf("%s/%s", c.host, uri)
}
