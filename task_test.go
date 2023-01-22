package main

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/account"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/api"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/task"
	"strconv"
	"testing"
)

func TestTask(t *testing.T) {
	api.Request.App = true
	api.Request.Cookie = ""
	if api.Request.Cookie == "" {
		t.Error("please set cookie")
	}
	accounts := account.GET_ACCOUNT_INFORMATION()
	Tasks := task.Task{AccountId: strconv.Itoa(accounts.Data.AccountID)}
	Tasks.NovelCompleteTas()
}
