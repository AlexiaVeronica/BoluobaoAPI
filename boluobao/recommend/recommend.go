package recommend

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/BoluobaoStructs"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/request"
)

func GET_SPECIALPUSH() BoluobaoStructs.Specialpush {
	var Specialpush BoluobaoStructs.Specialpush
	request.Get("specialpush").Add("expand", "novels").NewRequests().Unmarshal(&Specialpush)
	return Specialpush
}

func GET_ACTPUSH() BoluobaoStructs.Actpush {
	var Actpush BoluobaoStructs.Actpush
	params := map[string]string{"filter": "android", "page": "0", "size": "20"}
	request.Get("actpush").AddAll(params).NewRequests().Unmarshal(&Actpush)
	return Actpush
}

func GET_SPECIALPUSHS() BoluobaoStructs.FXrecommend {
	var FXrecommend BoluobaoStructs.FXrecommend
	params := map[string]string{"pushNames": "hotpush", "page": "0", "size": "10", "expand": "sysTags,discount,discountExpireDate,homeFlag"}
	request.Get("novels/specialpushs").AddAll(params).NewRequests().Unmarshal(&FXrecommend)
	return FXrecommend
}

func GET_SPECIALPUSHS2() BoluobaoStructs.NewBookRecommend {
	var NewBookRecommend BoluobaoStructs.NewBookRecommend
	params := map[string]string{"pushNames": "newpush", "page": "0", "size": "10", "expand": "sysTags,discount,discountExpireDate,homeFlag"}
	request.Get("novels/specialpushs").AddAll(params).NewRequests().Unmarshal(&NewBookRecommend).WriteResultString()
	return NewBookRecommend
}
