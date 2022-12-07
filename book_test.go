package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/rank"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/tag"
	"github.com/VeronicaAlexia/BoluobaoAPI/pkg/config"
	"testing"
)

func TestBook(t *testing.T) {
	config.AppConfig.App = true
	tag.PrintTagList()
	RESULT := tag.GET_TAG_INFO_ALL("", 100)
	fmt.Println(RESULT)
	fmt.Println(len(RESULT))

	var index int
	//view, sale, newhit, mark, ticket, bonus
	ranks := rank.Rank{RtypeIndex: 0, Page: 3}
	for _, book_info := range ranks.GET_SFACG_RANKS().Data {
		var SysTagsList []string
		index++
		SysTags := book_info.Expand.SysTags
		for _, value := range SysTags {
			SysTagsList = append(SysTagsList, value.TagName)
		}
		fmt.Println("index:", index, book_info.NovelName)
	}
}
