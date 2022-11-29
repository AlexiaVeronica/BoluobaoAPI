package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/book"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"os"
	"strconv"
)

func GetContent(ChapID string) {
	contents := book.GET_CHAPTER_CONTENT(ChapID)
	//fmt.Println(contents)
	if contents.Status.HTTPCode == 200 {
		content_text := []byte("\n\n\n" + contents.Data.Title + "\n\n" + contents.Data.Expand.Content)
		path := fmt.Sprintf("%v.txt", BookInfo.Data.NovelName)
		fl, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return
		}
		defer fl.Close()
		if _, err = fl.Write(content_text); err != nil {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("Content Error:", contents.Status.Msg)
	}
}
func GetChapter(book_id string) {
	for _, v := range book.GET_CATALOGUE(book_id).Data.VolumeList {
		fmt.Println("VolumeName:", v.Title)
		for _, v2 := range v.ChapterList {
			fmt.Println("	ChapterName:", v2.Title)
			GetContent(strconv.Itoa(v2.ChapID))
		}
	}
}

var BookInfo Template.BookInfo

func main() {
	book_id := "551946"
	App := request.AppRequest{App: false}
	App.SetApiHost()
	BookInfo = book.GET_BOOK_INFORMATION(book_id)
	if BookInfo.Status.HTTPCode == 200 {
		fmt.Println("bookName:", BookInfo.Data.NovelName)
		fmt.Println("AuthorName:", BookInfo.Data.AuthorName)
		fmt.Println("BookID:", BookInfo.Data.NovelID)
		fmt.Println("bookCover:", BookInfo.Data.NovelCover)

		if err := os.WriteFile(
			fmt.Sprintf("%v.txt", BookInfo.Data.NovelName),
			[]byte(BookInfo.Data.NovelName+"\n\n"), 0777); err != nil {
			fmt.Println(err)
		}
		GetChapter(strconv.Itoa(BookInfo.Data.NovelID))
	} else {
		fmt.Println("BookInfo Error:", BookInfo.Status.Msg)
	}
}
