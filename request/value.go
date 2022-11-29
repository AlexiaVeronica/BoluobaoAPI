package request

import (
	"bytes"
	"net/http"
	"os"
	"strings"
)

func (is *HttpUtils) GetEncodeParams() *bytes.Reader {
	return bytes.NewReader([]byte(is.query_data.Encode()))
}
func (is *HttpUtils) PutData() *strings.Reader {
	return strings.NewReader(is.DataFormString)
}
func (is *HttpUtils) GetResult() string {
	return string(is.result_body)
}

func (is *HttpUtils) WriteResultString() {
	_ = os.WriteFile("result.json", []byte(is.GetResult()), 0666)
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

func (is *HttpUtils) AddAll(params map[string]string) *HttpUtils {
	for key, value := range params {
		is.query_data.Add(key, value)
	}
	return is
}

func (is *HttpUtils) AddString(value string) *HttpUtils {
	is.DataFormString = value
	return is
}
