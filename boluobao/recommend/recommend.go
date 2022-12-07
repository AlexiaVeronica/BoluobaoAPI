package recommend

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
)

type Recommend struct {
	pushNames string
}

func (recommend *Recommend) GET_SPECIAL_PUSH() Template.Specialpush {
	var Specialpush Template.Specialpush
	request.Get("specialpush").Add("expand", "novels").NewRequests().Unmarshal(&Specialpush)
	return Specialpush
}

func (recommend *Recommend) GET_ACT_PUSH() Template.Actpush {
	var Actpush Template.Actpush
	params := map[string]string{"filter": "android", "page": "0", "size": "20"}
	request.Get("actpush").AddAll(params).NewRequests().Unmarshal(&Actpush)
	return Actpush
}

func (recommend *Recommend) GET_HOST_PUSH() Template.FXrecommend {
	var FXrecommend Template.FXrecommend
	params := map[string]string{"pushNames": "hotpush", "page": "0", "size": "10", "expand": "sysTags,homeFlag"}
	request.Get("novels/specialpushs").AddAll(params).NewRequests().Unmarshal(&FXrecommend)
	return FXrecommend
}

func (recommend *Recommend) GET_NEW_BOOK_PUSH() Template.NewBookRecommend {
	var NewBookRecommend Template.NewBookRecommend
	params := map[string]string{"pushNames": "newpush", "page": "0", "size": "10", "expand": "sysTags,homeFlag"}
	request.Get("novels/specialpushs").AddAll(params).NewRequests().Unmarshal(&NewBookRecommend)
	return NewBookRecommend
}
