package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao"
	"reflect"
	"testing"
)

func RangeStruct() []string {
	// 遍历rank.TypeName
	var typeName []string
	v := reflect.ValueOf(boluobao.API.Rank.TypeNameInit())
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
		boluobao.API.Rank.TypeName = value
		boluobao.API.Rank.Month = true
		boluobao.API.Rank.Page = 0
		for _, i2 := range boluobao.API.Rank.GET_SFACG_RANKS().Data {
			index++
			println(index, i2.NovelName)
			var a []string
			for _, tag := range i2.Expand.SysTags {
				a = append(a, tag.TagName)
			}
		}
	}

}
