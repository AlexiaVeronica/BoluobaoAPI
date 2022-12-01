package request

import (
	"encoding/json"
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/config"
	"io"
	"net/http"
	"net/url"
)

type HttpUtils struct {
	url            string
	method         string
	Cookie         string
	cookie         []*http.Cookie
	response       *http.Request
	query_data     *url.Values
	DataFormString string
	result_body    []byte
}

func NewHttpUtils(host string, path string, method string) *HttpUtils {
	return &HttpUtils{
		method:     method,
		query_data: &url.Values{},
		url:        host + path,
		Cookie:     config.AppConfig.Cookie,
	}
}

func (is *HttpUtils) NewRequests() *HttpUtils {
	var err error
	if is.method == "GET" {
		is.response, err = http.NewRequest(is.method, is.url+"?"+is.query_data.Encode(), nil)
	} else if is.method == "POST" {
		is.response, err = http.NewRequest(is.method, is.url, is.GetEncodeParams())
	} else if is.method == "PUT" {
		is.response, err = http.NewRequest(is.method, is.url, is.PutData())
	} else {
		panic("method error" + is.method)
	}
	if err != nil {
		panic(err)
	}
	is.result_body = nil
	is.set_headers()

	if response, ok := http.DefaultClient.Do(is.response); ok == nil {
		is.cookie = response.Cookies()
		is.result_body, _ = io.ReadAll(response.Body)
	} else {
		fmt.Println("NewRequests:", ok)
	}
	return is
}

func (is *HttpUtils) Unmarshal(s any) *HttpUtils {
	err := json.Unmarshal(is.result_body, s)
	if err != nil {
		fmt.Println("Unmarshal: ", err)
	}
	return is
}
