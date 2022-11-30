package tag

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"strconv"
)

func GET_SYS_TAG_LIST() Template.SysTags {
	var SysTags Template.SysTags
	request.Get("novels/0/sysTags").Add("filter", "push").NewRequests().Unmarshal(&SysTags)
	return SysTags
}

func GET_TAG_INFO(TagID string, page int) Template.Tag {
	params := map[string]string{
		"systagids":      TagID,
		"isfree":         "both",
		"size":           "50",
		"charcountbegin": "0",
		"updatedays":     "-1",
		"expand":         "typeName,sysTags,discount,discountExpireDate",
		"sort":           "latest",
		"page":           strconv.Itoa(page),
		"charcountend":   "0",
		"isfinish":       "both",
	}
	var BookTags Template.Tag
	request.Get("novels/0/sysTags/novels").AddAll(params).NewRequests().Unmarshal(&BookTags)
	return BookTags
}

func GET_TAG_INFO_ALL(TagID string, last_page int) []Template.Tag {
	var TagsList []Template.Tag
	for i := 1; i <= last_page; i++ {
		response := GET_TAG_INFO(TagID, last_page)
		if response.Status.HTTPCode == 200 {
			TagsList = append(TagsList, response)
		}
	}
	return TagsList
}
