package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type HttpUtils struct {
	url         string
	method      string
	Cookie      string
	cookie      []*http.Cookie
	response    *http.Request
	app_type    string
	query_data  *url.Values
	result_body []byte
}

func (is *HttpUtils) GetEncodeParams() *bytes.Reader {
	return bytes.NewReader([]byte(is.query_data.Encode()))
}
func (is *HttpUtils) GetResultBody() string {
	return string(is.result_body)
}

func (is *HttpUtils) GetCookie() []*http.Cookie {
	return is.cookie
}
func (is *HttpUtils) GetValue(key string) string {
	return is.query_data.Get(key)
}

func (is *HttpUtils) GetUrl() string {
	return is.url
}

func (is *HttpUtils) Add(key string, value string) *HttpUtils {
	is.query_data.Add(key, value)
	return is
}
func MustNewRequest(method, url string, data io.Reader) *http.Request {
	if request, err := http.NewRequest(method, url, data); err != nil {
		fmt.Println("Error: ", err)
		return nil
	} else {
		return request
	}
}

func Get(api_url string) *HttpUtils {
	return NewHttpUtils(api_url, "GET")
}
func Post(api_url string) *HttpUtils {
	return NewHttpUtils(api_url, "POST")
}

func NewHttpUtils(api_url, method string) *HttpUtils {
	return &HttpUtils{
		method:     method,
		query_data: &url.Values{},
		url:        "https://minipapi.sfacg.com/pas/mpapi/" + strings.ReplaceAll(api_url, "https://minipapi.sfacg.com/pas/mpapi/", ""),
	}
}

func (is *HttpUtils) NewRequests() *HttpUtils {
	is.result_body = nil
	is.response = MustNewRequest(is.method, is.url, is.GetEncodeParams())
	is.response.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	is.response.Header.Set("sf-minip-info", "minip_novel/1.0.70(android;11)/wxmp")
	is.response.Header.Set("Cookie", is.Cookie)

	if response, ok := http.DefaultClient.Do(is.response); ok == nil {
		is.cookie = response.Cookies()
		is.result_body, _ = io.ReadAll(response.Body)
	} else {
		fmt.Println("Error: ", ok)
	}
	return is
}

func (is *HttpUtils) Unmarshal(s any) *HttpUtils {
	err := json.Unmarshal(is.result_body, s)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return is
}
