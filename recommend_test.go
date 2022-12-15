package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/recommend"
	"testing"
)

func TestRecommend(t *testing.T) {
	Recommend := recommend.Recommend{}
	Recommend.GET_SPECIAL_PUSH()
	Recommend.GET_ACT_PUSH()
	fmt.Println("TestRecommend")
	for i, i2 := range Recommend.GET_HOST_PUSH().Data.HotPush {
		fmt.Println(i, ":", i2.Novel.NovelName)
	}
	fmt.Println("GET_HOST_PUSH Done")
	for i, i2 := range Recommend.GET_NEW_BOOK_PUSH().Data.NewPush {
		fmt.Println(i, ":", i2.Novel.NovelName)
	}

}
