package tag

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/liushuochen/gotable"
	"strconv"
	"sync"
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
	var TagsList []Template.BookInfoData
	var wg sync.WaitGroup
	var ch = make(chan interface{}, 64)
	for i := 1; i <= last_page; i++ {
		wg.Add(1)
		go func(page int) {
			ch <- "ok"
			defer wg.Done()
			response := GET_TAG_INFO(TagID, page)
			if response.Status.HTTPCode == 200 && response.Data != nil {
				for _, data := range response.Data {
					TagsList = append(TagsList, data)
				}
			} else {
				fmt.Println(response.Status.Msg)
			}
			<-ch
		}(i)
	}
	wg.Wait()
	fmt.Println("GET_TAG_INFO_ALL Done", len(TagsList))
	return TagsList
}
