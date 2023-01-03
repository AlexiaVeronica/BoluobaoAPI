package bookshelf

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
)

func NovelBookShelf() *[]Template.ShelfData {
	response := GET_BOOK_SHELF_INFORMATION()
	if response.Status.HTTPCode == 200 {
		return &response.Data
	}
	fmt.Println("BookShelf Error:", response.Status.Msg)
	return nil
}
