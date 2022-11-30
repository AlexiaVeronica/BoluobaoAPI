package tag

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"sync"
)

func PrintTagList() {
	for _, value := range GET_SYS_TAG_LIST().Data {
		fmt.Println("tag_id:", value.SysTagID, "\t\ttag_name:", value.TagName)
	}
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
