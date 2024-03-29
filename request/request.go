package request

import (
	"encoding/json"
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
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
	ResultBody     []byte
	Config         Template.ConfigRequest
}

func NewHttpUtils(host string, path string, method string, Config Template.ConfigRequest) *HttpUtils {
	return &HttpUtils{
		method:     method,
		query_data: &url.Values{},
		url:        host + path,
		Cookie:     Config.Cookie,
		Config:     Config,
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
	is.ResultBody = nil
	is.SetHeaders()

	if response, ok := http.DefaultClient.Do(is.response); ok == nil {
		is.cookie = response.Cookies()
		is.ResultBody, _ = io.ReadAll(response.Body)
	} else {
		fmt.Println("NewRequests:", ok)
	}
	return is
}

func (is *HttpUtils) Unmarshal(s any) *HttpUtils {
	err := json.Unmarshal(is.ResultBody, s)
	if err != nil {
		fmt.Println("Unmarshal: ", err)
	}
	return is
}
