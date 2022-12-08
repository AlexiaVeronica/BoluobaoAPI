package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/account"
	"github.com/VeronicaAlexia/BoluobaoAPI/pkg/config"
	"testing"
)

func TestBook(t *testing.T) {
	config.AppConfig.App = false
	fmt.Println(account.GET_USER_WORKS("").Data)

}
