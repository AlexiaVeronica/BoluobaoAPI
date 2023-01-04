package recommend

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"strings"
)

func SysTags(tags []Template.SysTagList) string {
	var tagString []string
	for _, tag := range tags {
		tagString = append(tagString, tag.TagName)
	}
	return strings.Join(tagString, ",")
}

func (recommend *Recommend) NovelRecommend() {
	response := recommend.GET_HOST_PUSH()
	for index, data := range response.Data.HotPush {
		fmt.Println("index:", index, "\tbookName:", data.Novel.NovelName, "\tTag:", SysTags(data.Novel.Expand.SysTags))
	}
	response2 := recommend.GET_NEW_BOOK_PUSH()
	for index, data := range response2.Data.NewPush {
		fmt.Println("index:", index, "\tbookName:", data.Novel.NovelName, "\tTag:", SysTags(data.Novel.Expand.SysTags))
	}
}
