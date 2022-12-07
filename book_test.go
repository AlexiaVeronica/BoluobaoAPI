package main

import (
	"fmt"
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

}
