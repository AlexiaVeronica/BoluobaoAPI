package request

import (
	"strings"
)

func (is *HttpUtils) set_headers() {
	Header := make(map[string]string)
	Header["Connection"] = "keep-alive"
	if is.method == "POST" {
		Header["Content-Type"] = "application/x-www-form-urlencoded"
	} else {
		Header["Content-Type"] = "application/json"
	}
	Header["Cookie"] = is.Cookie
	if strings.Contains(is.url, "api.sfacg.com") {
		Header["Host"] = "api.sfacg.com"
		Header["User-Agent"] = "boluobao/4.8.42(android;25)/XIAOMI/240a90cc-4c40-32c7-b44e-d4cf9e670605/XIAOMI"
		Header["authorization"] = "Basic YW5kcm9pZHVzZXI6MWEjJDUxLXl0Njk7KkFjdkBxeHE="
	} else {
		Header["sf-minip-info"] = "minip_novel/1.0.70(android;11)/wxmp"
	}
	//fmt.Println(Header)
	for HeaderKey, HeaderValue := range Header {
		is.response.Header.Set(HeaderKey, HeaderValue)

	}
}
