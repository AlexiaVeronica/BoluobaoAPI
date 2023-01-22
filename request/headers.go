package request

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (is *HttpUtils) WeChatBasic() string {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	signMd5 := md5.New()
	signMd5.Write([]byte("d3htaW5pYXBw" + timestamp + "null" + "xw3#a12-x"))
	sign := strings.ToUpper(hex.EncodeToString(signMd5.Sum(nil)))
	wxmpuser := "wxmpuser:194c5#b_47Fc75676f:nonce=%v&deviceToken=null&timestamp=%v&sign=%v"
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(wxmpuser, "d3htaW5pYXBw", timestamp, sign)))
}

func (is *HttpUtils) SetHeaders() {
	Header := make(map[string]string)
	Header["Connection"] = "keep-alive"
	if is.method == "POST" {
		Header["Content-Type"] = "application/x-www-form-urlencoded"
	} else {
		Header["Content-Type"] = "application/json"
	}
	Header["Cookie"] = is.Cookie
	if strings.Contains(is.url, "minipapi.sfacg.com") {
		Header["sf-minip-info"] = "minip_novel/1.0.70(android;11)/wxmp"
		Header["Authorization"] = "Basic " + is.WeChatBasic()
	} else {
		Header["Host"] = "api.sfacg.com"
		Header["User-Agent"] = "boluobao/4.8.42(android;25)/XIAOMI/240a90cc-4c40-32c7-b44e-d4cf9e670605/XIAOMI"
		Header["Authorization"] = "Basic YW5kcm9pZHVzZXI6MWEjJDUxLXl0Njk7KkFjdkBxeHE="
		Header["SFSecurity"] = is.AppSecurity()

	}
	for HeaderKey, HeaderValue := range Header {
		is.response.Header.Set(HeaderKey, HeaderValue)

	}
}

func (is *HttpUtils) AppSecurity() string {
	TimeStamp := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	uuId, newMd5 := strings.ToUpper(uuid.New().String()), md5.New()
	newMd5.Write([]byte(uuId + TimeStamp + strings.ToUpper(is.Config.DeviceId) + is.Config.AppKey))
	security_params := url.Values{
		"nonce":       {uuId},
		"timestamp":   {TimeStamp},
		"devicetoken": {strings.ToUpper(is.Config.DeviceId)},
		"sign":        {strings.ToUpper(hex.EncodeToString(newMd5.Sum(nil)))},
	}
	return security_params.Encode()
}
