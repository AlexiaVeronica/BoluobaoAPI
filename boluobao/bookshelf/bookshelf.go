package bookshelf

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
)

type BookShelf struct{}

func (bs *BookShelf) NovelBookShelf() *[]Template.ShelfData {
	response := GET_BOOK_SHELF_INFORMATION()
	if response.Status.HTTPCode == 200 {
		return &response.Data
	}
	fmt.Println("BookShelf Error:", response.Status.Msg)
	return nil
}
