# BoluobaoApi

## example

```
package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/book"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/config"
	"os"
	"testing"
)

func GetContent(bookInfo *Template.BookInfoData, ChapID string) {
	contents := book.NovelContent(ChapID)
	if contents != nil {
		content_text := []byte("\n\n\n" + contents.Data.Expand.Content)
		path := fmt.Sprintf("%v.txt", bookInfo.NovelName)
		fl, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return
		}
		defer fl.Close()
		if _, err = fl.Write(content_text); err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func main() {
	book_id := "512854"
	api.Request.App = true
	api.Request.DeviceId = ""
	api.Request.AppKey = ""

	bookInfo := book.NovelInfo(book_id)
	if bookInfo != nil {
		fmt.Println("bookName:", bookInfo.NovelName)
		fmt.Println("AuthorName:", bookInfo.AuthorName)
		fmt.Println("BookID:", bookInfo.NovelID)
		fmt.Println("bookCover:", bookInfo.NovelCover)

		if err := os.WriteFile(
			fmt.Sprintf("%v.txt", bookInfo.NovelName),
			[]byte(bookInfo.NovelName+"\n\n"), 0777); err != nil {
			fmt.Println(err)
		}
		for _, ChapID := range book.NovelCatalogue(book_id) {
			GetContent(bookInfo, ChapID)
		}
	}
}


```