package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/account"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"testing"
)

func TestBook(t *testing.T) {
	AccountId := ""
	request.AppConfig.App = false
	fmt.Println(account.GET_USER_COMMENT(AccountId, 0))
	account.LOGIN_ACCOUNT("username", "password")
	account.GET_ACCOUNT_INFORMATION()
	//account.GET_ACCOUNT_MONEY_INFORMATION()
	account.GET_USER_INFORMATION(AccountId)
	fmt.Println(account.GET_USER_WORKS(AccountId).Data)
	account.GET_ACCOUNT_IP()

}
