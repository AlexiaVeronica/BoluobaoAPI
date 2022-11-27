# BoluobaoApi

## example

```go
package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/BoluobaoStructs"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/request"
	"os"
	"strconv"
	"testing"
)

func GetContent(ChapID string) {
	contents := boluobao.GET_CHAPTER_CONTENT(ChapID)
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
	}
}
func GetChapter(book_id string) {
	for _, v := range boluobao.GET_CATALOGUE(book_id).Data.VolumeList {
		fmt.Println("VolumeName:", v.Title)
		for _, v2 := range v.ChapterList {
			fmt.Println("	ChapterName:", v2.Title)
			GetContent(strconv.Itoa(v2.ChapID))
		}
	}
}

var BookInfo BoluobaoStructs.BookInfo

func TestDownload(t *testing.T) {
	book_id := "551946"
	request.APP_TYPE.App = "wx"
	request.APP_TYPE.Host = "https://minipapi.sfacg.com/pas/mpapi/"
	BookInfo = boluobao.GET_BOOK_INFORMATION(book_id)
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
		fmt.Println("Error:", BookInfo.Status.Msg)
	}
}

```