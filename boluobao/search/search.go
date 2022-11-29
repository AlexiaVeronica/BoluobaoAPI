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
	var wg = sync.WaitGroup{}
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			response := GET_SEARCH(keyword, page)
			if response.Status.HTTPCode == 200 {
				SearchList = append(SearchList, response)
			} else {
				fmt.Println("search failed:", response.Status.Msg)
			}
		}(i)
	}
	wg.Wait() // wait for all goroutines to finish
	return SearchList
}
