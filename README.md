# BoluobaoApi

## example

```go
package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/BoluobaoStructs"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao"
	"os"
	"strconv"
)

func GetContent(ChapID string) {
	contents := boluobao.GET_CHAPTER_CONTENT(ChapID)
	if contents.Status.HTTPCode == 200 {
		content_text := []byte(contents.Data.Expand.Content)
		path := fmt.Sprintf("%v.txt", BookInfo.Data.NovelName)
		if err := os.WriteFile(path, content_text, 0666); err != nil {
			fmt.Println(err)
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

func main() {
	book_id := "551946"
	BookInfo = boluobao.GET_BOOK_INFORMATION(book_id)
	if BookInfo.Status.HTTPCode == 200 {
		fmt.Println("bookName:", BookInfo.Data.NovelName)
		fmt.Println("AuthorName:", BookInfo.Data.AuthorName)
		fmt.Println("BookID:", BookInfo.Data.NovelID)
		fmt.Println("bookCover:", BookInfo.Data.NovelCover)
		GetChapter(strconv.Itoa(BookInfo.Data.NovelID))
	} else {
		fmt.Println("Error:", BookInfo.Status.Msg)
	}
}
```