package main

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/account"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/task"
	"github.com/VeronicaAlexia/BoluobaoAPI/config"
	"strconv"
	"testing"
)

func TestTask(t *testing.T) {
	config.AppConfig.App = true
	config.AppConfig.Cookie = ""
	accounts := account.GET_ACCOUNT_INFORMATION()
	Tasks := task.Task{AccountId: strconv.Itoa(accounts.Data.AccountID)}
	Tasks.COMPLETE_ALL()
}
