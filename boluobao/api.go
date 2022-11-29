package boluobao

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"strconv"
)

func GET_BOOK_SHELF_INFORMATION() Template.InfoData {
	var bookshelfData Template.InfoData
	request.Get("user/Pockets").Add("expand", "novels").NewRequests().Unmarshal(&bookshelfData)
	return bookshelfData
}

func GET_SYS_TAG_LIST() Template.SysTags {
	var SysTags Template.SysTags
	request.Get("novels/0/sysTags").Add("filter", "push").NewRequests().Unmarshal(&SysTags)
	return SysTags
}

func GET_SEARCH(keyword string, page int) Template.Search {
	var Search Template.Search
	request.Get("search/novels/result").Add("q", keyword).
		Add("size", "20").Add("page", strconv.Itoa(page)).NewRequests().Unmarshal(&Search)
	return Search
}
