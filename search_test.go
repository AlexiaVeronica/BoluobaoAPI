package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/search"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"testing"
)

func TestSearch(t *testing.T) {
	App := request.AppRequest{App: true}
	App.SetApiHost()

	response := search.GET_SEARCH("斗罗大陆", 0)
	if response.Status.HTTPCode != 200 {
		fmt.Println(response.Status.Msg, "search failed")
	} else {
		for i, data := range response.Data.Novels {
			fmt.Println("index:", i, "\tbookID:", data.NovelID, "\tbookName:", data.NovelName)
		}
	}
}

func TestSearchAll(t *testing.T) {
	App := request.AppRequest{App: true}
	App.SetApiHost()
	response_list := search.GET_SEARCH_All("血姬")
	var index int
	for _, response := range response_list {
		for _, data := range response.Data.Novels {
			index++
			fmt.Println("index:", index, "\tbookID:", data.NovelID, "\tbookName:", data.NovelName)
		}
	}
	fmt.Println(len(response_list))
}
