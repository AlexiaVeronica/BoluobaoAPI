package request

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func WeChatBasic() string {
	signString := "d3htaW5pYXBw" + strconv.FormatInt(time.Now().UnixNano()/1000000, 10) + "null" + "xw3#a12-x"
	signMd5 := md5.New()
	signMd5.Write([]byte(signString))
	Authorization := fmt.Sprintf("wxmpuser:194c5#b_47Fc75676f:nonce=%v&deviceToken=null&timestamp=%v&sign=%v",
		"d3htaW5pYXBw", strconv.FormatInt(time.Now().UnixNano()/1000000, 10), strings.ToUpper(hex.EncodeToString(signMd5.Sum(nil))))
	return base64.StdEncoding.EncodeToString([]byte(Authorization))
}

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
		Header["authorization"] = "Basic " + WeChatBasic()
	}
	//fmt.Println(Header)
	for HeaderKey, HeaderValue := range Header {
		is.response.Header.Set(HeaderKey, HeaderValue)

	}
}
