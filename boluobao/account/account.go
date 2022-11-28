package account

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
)

func GET_ACCOUNT_INFORMATION() Template.Account {
	var Account Template.Account
	request.Get("user").NewRequests().Unmarshal(&Account)
	return Account
}

func LOGIN_ACCOUNT(username, password string) string {
	var Cookie string
	params := map[string]string{"username": username, "password": password}
	CookieJar := request.Post("sessions").AddAll(params).NewRequests().Unmarshal(&Template.Login).GetCookie()
	if Template.Login.Status.HTTPCode == 200 {
		for _, cookie := range CookieJar {
			Cookie += cookie.Name + "=" + cookie.Value + ";"
		}
		return Cookie
	} else {
		fmt.Println("login failed! error:", Template.Login.Status.Msg)
		return ""
	}
}
