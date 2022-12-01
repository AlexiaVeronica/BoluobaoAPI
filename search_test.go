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
	for i, data := range response.Data.Novels {
		fmt.Println("index:", i, "\tbookID:", data.NovelID, "\tbookName:", data.NovelName)
	}

}

func TestSearchAll(t *testing.T) {
	App := request.AppRequest{App: true}
	App.SetApiHost()
	response_list := search.GET_SEARCH_All("血姬")

	if len(response_list) == 0 {
		fmt.Println("没有搜索到结果,请检查关键词是否正确")
	} else {
		fmt.Println("搜索完成:", len(response_list), "页")
		for index, data := range response_list {
			fmt.Println("index:", index, "\tbookID:", data.NovelID, "\tbookName:", data.NovelName)
		}
	}
}
