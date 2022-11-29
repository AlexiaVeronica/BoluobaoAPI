package main

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/account"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/task"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"strconv"
	"testing"
)

func TestTask(t *testing.T) {
	App := request.AppRequest{App: true}
	App.SetApiHost()
	request.Cookie = ""
	accounts := account.GET_ACCOUNT_INFORMATION()
	Tasks := task.Task{AccountId: strconv.Itoa(accounts.Data.AccountID)}
	Tasks.COMPLETE_ALL()
}
