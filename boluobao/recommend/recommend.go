package recommend

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
)

func GET_SPECIALPUSH() Template.Specialpush {
	var Specialpush Template.Specialpush
	request.Get("specialpush").Add("expand", "novels").NewRequests().Unmarshal(&Specialpush)
	return Specialpush
}

func GET_ACTPUSH() Template.Actpush {
	var Actpush Template.Actpush
	params := map[string]string{"filter": "android", "page": "0", "size": "20"}
	request.Get("actpush").AddAll(params).NewRequests().Unmarshal(&Actpush)
	return Actpush
}

func GET_SPECIALPUSHS() Template.FXrecommend {
	var FXrecommend Template.FXrecommend
	params := map[string]string{"pushNames": "hotpush", "page": "0", "size": "10", "expand": "sysTags,discount,discountExpireDate,homeFlag"}
	request.Get("novels/specialpushs").AddAll(params).NewRequests().Unmarshal(&FXrecommend)
	return FXrecommend
}

func GET_SPECIALPUSHS2() Template.NewBookRecommend {
	var NewBookRecommend Template.NewBookRecommend
	params := map[string]string{"pushNames": "newpush", "page": "0", "size": "10", "expand": "sysTags,discount,discountExpireDate,homeFlag"}
	request.Get("novels/specialpushs").AddAll(params).NewRequests().Unmarshal(&NewBookRecommend).WriteResultString()
	return NewBookRecommend
}
