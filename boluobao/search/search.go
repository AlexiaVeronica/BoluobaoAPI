package search

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/pkg/threading"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"strconv"
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

func GET_SEARCH_All(keyword string, lastPage int) []Template.BookInfoData {
	var BookInfoList []Template.BookInfoData
	Thread := threading.InitThreading(64)
	for i := 0; i < lastPage; i++ {
		Thread.Add()
		go func(page int) {
			defer Thread.Done()
			response := GET_SEARCH(keyword, page)
			if response != nil {
				for _, data := range response.Data.Novels {
					BookInfoList = append(BookInfoList, data)
				}
			}
		}(i)
	}
	Thread.Wait() // wait for all goroutines to finish
	return BookInfoList
}
