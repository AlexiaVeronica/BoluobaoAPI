package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/account"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/api"
	"testing"
)

func TestBook(t *testing.T) {
	AccountId := ""
	api.Request.App = false
	fmt.Println(account.GET_USER_COMMENT(AccountId, 0))
	cookie := account.LOGIN_ACCOUNT("username", "password")
	if cookie != "" {
		fmt.Println("登录成功")
		api.Request.Cookie = cookie
	} else {
		fmt.Println("登录失败")
	}
	if api.Request.Cookie == "" {
		t.Error("please set cookie")
	}
	account.GET_ACCOUNT_INFORMATION()
	//account.GET_ACCOUNT_MONEY_INFORMATION()
	account.GET_USER_INFORMATION(AccountId)
	fmt.Println(account.GET_USER_WORKS(AccountId).Data)
	account.GET_ACCOUNT_IP()

}
