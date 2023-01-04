package search

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/pkg/threading"
)

type Search struct{}

func (s *Search) NovelSearch(keyword string, lastPage int) []Template.BookInfoData {
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
