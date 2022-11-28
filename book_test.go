package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/rank"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"

	"testing"
)

func TestBook(t *testing.T) {
	App := request.AppRequest{App: false}
	App.SetApiHost()
	ranks := rank.Rank{RtypeIndex: 0, Page: 0}
	for index, i2 := range ranks.GET_SFACG_RANKS().Data {
		fmt.Println(index, ": ", i2.NovelName)
	}
}

//func TestCATALOGUE(t *testing.T) {
//	fmt.Println(boluobao.GET_CATALOGUE("551946"))
//}
