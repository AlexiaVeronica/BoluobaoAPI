package request

import "fmt"

func (is *HttpUtils) set_headers() {
	Header := make(map[string]string)
	Header["Connection"] = "keep-alive"
	if is.method == "POST" {
		Header["Content-Type"] = "application/x-www-form-urlencoded"
	} else {
		Header["Content-Type"] = "application/json"
	}
	Header["Cookie"] = is.Cookie
	if APP_TYPE.App == "app" {
		Header["Host"] = "api.sfacg.com"
		Header["User-Agent"] = "boluobao/4.8.42(android;25)/XIAOMI/240a90cc-4c40-32c7-b44e-d4cf9e670605/XIAOMI"
		Header["authorization"] = "Basic YW5kcm9pZHVzZXI6MWEjJDUxLXl0Njk7KkFjdkBxeHE="
	} else if APP_TYPE.App == "wx" {
		Header["sf-minip-info"] = "minip_novel/1.0.70(android;11)/wxmp"
	} else {
		fmt.Println("Error: APP_TYPE.App is not set")
	}
	//fmt.Println(Header)
	for HeaderKey, HeaderValue := range Header {
		is.response.Header.Set(HeaderKey, HeaderValue)

	}
}
