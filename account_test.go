package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/account"
	"github.com/VeronicaAlexia/BoluobaoAPI/pkg/config"
	"testing"
)

func TestBook(t *testing.T) {
	config.AppConfig.App = false
	account.GET_ACCOUNT_IP()
	fmt.Println(account.GET_USER_WORKS("").Data)

}
