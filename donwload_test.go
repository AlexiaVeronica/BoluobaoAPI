package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/api"
	"os"
	"testing"
)

func GetContent(bookInfo *Template.BookInfoData, ChapID string) {
	contents := boluobao.API.Book.NovelContent(ChapID)
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

func TestDownload(t *testing.T) {
	book_id := "512854"
	api.Request.App = true
	api.Request.DeviceId = "240a90cc-4c40-32c7-b44e-d4cf9e670605"
	api.Request.AppKey = ""
	api.Request.Cookie = ""
	if api.Request.Cookie == "" {
		fmt.Println("you need to set cookie")
	}
	bookInfo := boluobao.API.Book.NovelInfo(book_id)
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
		for _, ChapID := range boluobao.API.Book.NovelCatalogue(book_id) {
			GetContent(bookInfo, ChapID)
		}
	}
}
