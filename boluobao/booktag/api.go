package booktag

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/api"
)

func GET_TAG_INFO(TagID string, page int) Template.Tag {
	var BookTags Template.Tag
	api.Get("novels/0/sysTags/novels").AddAll(TagsParams(TagID, page, "50")).NewRequests().Unmarshal(&BookTags)
	return BookTags
}

func GET_SYS_TAG_LIST() Template.SysTags {
	var SysTags Template.SysTags
	api.Get("novels/0/sysTags").Add("filter", "push").NewRequests().Unmarshal(&SysTags)
	return SysTags
}
