package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/rank"
	"reflect"
	"testing"
)

func RangeStruct() []string {
	// 遍历rank.TypeName
	var typeName []string
	v := reflect.ValueOf(rank.TypeName)
	count := v.NumField()
	for i := 0; i < count; i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.String:
			typeName = append(typeName, f.String())
		case reflect.Int:
			fmt.Println(f.Int())
		}
	}
	return typeName
}
func TestRanks(t *testing.T) {
	index := 0
	for _, value := range RangeStruct() {
		ranks := rank.Rank{TypeName: value, Month: true, Page: 0}
		for _, i2 := range ranks.GET_SFACG_RANKS().Data {
			index++
			println(index, i2.NovelName)
			var a []string
			for _, tag := range i2.Expand.SysTags {
				a = append(a, tag.TagName)
			}
		}
	}

}
