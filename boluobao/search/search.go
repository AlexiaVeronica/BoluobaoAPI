package search

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"strconv"
	"sync"
)

func GET_SEARCH(keyword string, page int) Template.Search {
	var Search Template.Search
	request.Get("search/novels/result").Add("q", keyword).
		Add("size", "20").Add("page", strconv.Itoa(page)).NewRequests().Unmarshal(&Search)
	return Search
}

func GET_SEARCH_All(keyword string) []Template.Search {
	var SearchList []Template.Search
	var Search Template.Search
	var wg = sync.WaitGroup{}
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			request.Get("search/novels/result").Add("q", keyword).
				Add("size", "50").Add("page", strconv.Itoa(page)).NewRequests().Unmarshal(&Search)
			if Search.Status.HTTPCode == 200 {
				SearchList = append(SearchList, Search)
			} else {
				fmt.Println("search failed:", Search.Status.Msg)
			}
		}(i)
	}
	wg.Wait()
	if len(SearchList) == 0 {
		fmt.Println("没有搜索到结果,请检查关键词是否正确")
		return nil
	}
	fmt.Println("搜索完成")
	return SearchList
}
