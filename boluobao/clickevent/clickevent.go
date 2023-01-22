package clickevent

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/api"
)

func POST_POCKETS(novelId string) Template.Account {
	var Account Template.Account
	params := `{"novelId":` + novelId + `,"categoryId":0}`
	api.Post("pockets/-1/novels").NewRequests().AddString(params).Unmarshal(&Account).WriteResultString()
	return Account
}

func POST_FAVS(novelId string) Template.Account {
	var Account Template.Account
	api.Post("novels/" + novelId + "/favs").NewRequests().Unmarshal(&Account).WriteResultString()
	return Account
}
