package bookshelf

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/api"
)

func GET_BOOK_SHELF_INFORMATION() Template.InfoData {
	var bookshelfData Template.InfoData
	api.Get("user/Pockets").Add("expand", "novels").NewRequests().Unmarshal(&bookshelfData)
	return bookshelfData
}
