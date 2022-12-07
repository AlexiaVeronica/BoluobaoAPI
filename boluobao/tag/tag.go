package tag

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/config"
	"github.com/liushuochen/gotable"
	"strconv"
)

func PrintTagList() {
	table, err := gotable.Create("tag_name", "tag_id")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, value := range GET_SYS_TAG_LIST().Data {
		err = table.AddRow([]string{value.TagName, strconv.Itoa(value.SysTagID)})
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println(table)
}
func GET_TAG_INFO_ALL(TagID string, last_page int) []Template.BookInfoData {
	Thread := config.InitThreading(64)
	var TagsList []Template.BookInfoData
	for i := 1; i <= last_page; i++ {
		Thread.Add()
		go func(page int) {
			defer Thread.Done()
			response := GET_TAG_INFO(TagID, page)
			if response.Status.HTTPCode == 200 && response.Data != nil {
				for _, data := range response.Data {
					TagsList = append(TagsList, data)
				}
			} else {
				fmt.Println(response.Status.Msg)
			}
		}(i)
	}
	Thread.Wait()
	fmt.Println("GET_TAG_INFO_ALL Done", len(TagsList))
	return TagsList
}
