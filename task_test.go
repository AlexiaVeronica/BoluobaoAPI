package main

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/account"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/task"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"strconv"
	"testing"
)

func TestTask(t *testing.T) {
	request.AppConfig.App = true
	request.AppConfig.Cookie = ""
	accounts := account.GET_ACCOUNT_INFORMATION()
	Tasks := task.Task{AccountId: strconv.Itoa(accounts.Data.AccountID)}
	Tasks.NovelCompleteTas()
}
