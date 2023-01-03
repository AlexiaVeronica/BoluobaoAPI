package book

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"strconv"
	"strings"
)

func NovelInfo(book_id string) *Template.BookInfoData {
	if response := GET_BOOK_INFORMATION(book_id); response.Status.HTTPCode == 200 {
		return &response.Data
	} else {
		fmt.Println("BookInfo Error:", response.Status.Msg)
	}
	return nil
}

func Catalogue(book_id string) []string {
	var chapter_id_list []string
	response := GET_CATALOGUE(book_id)
	if response.Status.HTTPCode != 200 {
		fmt.Println("Catalog Error:", response.Status.Msg)
		return nil
	}

	for index, volume := range response.Data.VolumeList {
		fmt.Println("VolumeIndex:", index, "\tVolumeName:", volume.Title)
		for _, chapter := range volume.ChapterList {
			fmt.Println("	ChapterIndex:", chapter.ChapOrder, "\tChapterName:", chapter.Title)
			if chapter.OriginNeedFireMoney == 0 {
				chapter_id_list = append(chapter_id_list, strconv.Itoa(chapter.ChapID))
			}
		}
	}
	return chapter_id_list
}

func Content(chapter_id string) *Template.Content {
	var ContentText string
	response := GET_CHAPTER_CONTENT(chapter_id)
	if response.Status.HTTPCode == 200 {
		for _, content := range strings.Split(response.Data.Expand.Content, "\n") {
			if content != "" {
				ContentText += "　　" + content + "\n"
			}
		}
		response.Data.Expand.Content = response.Data.Title + "\n\n" + ContentText
		return &response
	} else {
		fmt.Println("Content Error:", response.Status.Msg)
	}
	return nil
}
