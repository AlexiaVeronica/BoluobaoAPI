package main

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/rank"
	"testing"
)

func TestRanks(t *testing.T) {
	index := 0
	ranks := rank.Rank{TypeName: rank.TypeName.NewRank, All: true, Page: 0}
	for _, i2 := range ranks.GET_SFACG_RANKS().Data {
		index++
		println(index, i2.NovelName)
		var a []string
		for _, tag := range i2.Expand.SysTags {
			a = append(a, tag.TagName)
		}
	}

}
