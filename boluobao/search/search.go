package search

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"strconv"
	"sync"
)

func GET_SEARCH(keyword string, page int) *Template.Search {
	var Search Template.Search
	params := map[string]string{"q": keyword, "page": strconv.Itoa(page), "size": "50"}
	request.Get("search/novels/result").AddAll(params).NewRequests().Unmarshal(&Search)
	if Search.Status.HTTPCode == 200 && Search.Data.Novels != nil {
		return &Search
	} else {
		fmt.Println("search failed:", Search.Status.Msg)
	}
	return nil
}

func GET_SEARCH_All(keyword string) []Template.BookInfoData {
	var (
		BookInfoList []Template.BookInfoData
		wg           sync.WaitGroup
	)
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			response := GET_SEARCH(keyword, page)
			if response != nil {
				for _, data := range response.Data.Novels {
					BookInfoList = append(BookInfoList, data)
				}
			}
		}(i)
	}
	wg.Wait() // wait for all goroutines to finish
	return BookInfoList
}
