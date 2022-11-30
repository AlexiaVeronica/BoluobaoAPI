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
		table.AddRow([]string{value.TagName, strconv.Itoa(value.SysTagID)})
	}
	fmt.Println(table)
}
func GET_TAG_INFO_ALL(TagID string, last_page int) []Template.Tag {
	var TagsList []Template.Tag
	var wg *sync.WaitGroup
	for i := 1; i <= last_page; i++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			response := GET_TAG_INFO(TagID, page)
			if response.Status.HTTPCode == 200 {
				TagsList = append(TagsList, response)
			}
		}(i)
	}
	wg.Wait()
	return TagsList
}
