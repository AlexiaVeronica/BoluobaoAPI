package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/search"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"testing"
)

func TestSearch(t *testing.T) {
	request.AppConfig.App = true
	response := search.GET_SEARCH("斗罗大陆", 0)
	for i, data := range response.Data.Novels {
		fmt.Println("index:", i, "\tbookID:", data.NovelID, "\tbookName:", data.NovelName)
	}

}

func TestSearchAll(t *testing.T) {
	request.AppConfig.App = true
	response_list := boluobao.API.Search.NovelSearch("血姬", 2)
	if len(response_list) == 0 {
		fmt.Println("没有搜索到结果,请检查关键词是否正确")
	} else {
		fmt.Println("搜索完成:", len(response_list), "页")
		for _, data := range response_list {
			fmt.Println(data.NovelName)
			//fmt.Println("index:", index, "\tbookID:", data.NovelID, "\tbookName:", data.NovelName)
		}
	}
}
