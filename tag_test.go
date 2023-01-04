package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao"
	"testing"
)

func TestTag(t *testing.T) {
	//tag.PrintTagList()
	for i, i2 := range boluobao.API.Tags.NovelTags("48", 100) {
		//fmt.Println(i, i2.NovelName)
		fmt.Println("index:", i, "\tbookID:", i2.NovelID, "\tbookName:", i2.NovelName)
	}

}
