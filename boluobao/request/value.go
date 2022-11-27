package request

import (
	"bytes"
	"net/http"
)

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
