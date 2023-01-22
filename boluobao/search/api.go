package search

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/api"
	"strconv"
)

func GET_SEARCH(keyword string, page int) *Template.Search {
	var Search Template.Search
	params := map[string]string{"q": keyword, "page": strconv.Itoa(page), "size": "50"}
	api.Get("search/novels/result").AddAll(params).NewRequests().Unmarshal(&Search)
	if Search.Status.HTTPCode == 200 && Search.Data.Novels != nil {
		return &Search
	} else {
		fmt.Println("search failed:", Search.Status.Msg)
	}
	return nil
}
