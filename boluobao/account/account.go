package account

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/BoluobaoStructs"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/request"
)

func GET_ACCOUNT_INFORMATION() BoluobaoStructs.Account {
	var Account BoluobaoStructs.Account
	request.Get("user").NewRequests().Unmarshal(&Account)
	return Account
}

func LOGIN_ACCOUNT(username, password string) string {
	var Cookie string
	params := map[string]string{"username": username, "password": password}
	CookieJar := request.Post("sessions").AddAll(params).NewRequests().Unmarshal(&BoluobaoStructs.Login).GetCookie()
	if BoluobaoStructs.Login.Status.HTTPCode == 200 {
		for _, cookie := range CookieJar {
			Cookie += cookie.Name + "=" + cookie.Value + ";"
		}
		return Cookie
	} else {
		fmt.Println("login failed! error:", BoluobaoStructs.Login.Status.Msg)
		return ""
	}
}
