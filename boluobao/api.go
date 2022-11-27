package boluobao

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/BoluobaoStructs"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/request"
	"strconv"
)

func GET_BOOK_INFORMATION(NovelId string) BoluobaoStructs.BookInfo {
	var BookInfo BoluobaoStructs.BookInfo
	expand := "chapterCount,bigBgBanner,bigNovelCover,typeName,intro,fav,ticket,pointCount,tags,sysTags,signlevel,discount,discountExpireDate,totalNeedFireMoney,rankinglist,originTotalNeedFireMoney,firstchapter,latestchapter,latestcommentdate,essaytag,auditCover,preOrderInfo,customTag,topic,unauditedCustomtag,homeFlag,isbranch"
	request.Get("novels/"+NovelId).Add("expand", expand).NewRequests().Unmarshal(&BookInfo)
	return BookInfo
}

func GET_ACCOUNT_INFORMATION() BoluobaoStructs.Account {
	var Account BoluobaoStructs.Account
	request.Get("user").NewRequests().Unmarshal(&Account)
	return Account
}

func GET_BOOK_SHELF_INFORMATION() BoluobaoStructs.InfoData {
	var bookshelfData BoluobaoStructs.InfoData
	request.Get("user/Pockets").Add("expand", "novels").NewRequests().Unmarshal(&bookshelfData)
	return bookshelfData
}

func GET_SPECIALPUSH() BoluobaoStructs.Specialpush {
	var Specialpush BoluobaoStructs.Specialpush
	request.Get("specialpush").Add("expand", "novels").NewRequests().Unmarshal(&Specialpush)
	return Specialpush
}

func GET_SYS_TAG_LIST() BoluobaoStructs.SysTags {
	var SysTags BoluobaoStructs.SysTags
	request.Get("novels/0/sysTags").Add("filter", "push").NewRequests().Unmarshal(&SysTags)
	return SysTags
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

func GET_CATALOGUE(NovelID string) BoluobaoStructs.Catalogue {
	var Catalogue BoluobaoStructs.Catalogue
	request.Get(fmt.Sprintf("novels/%v/dirs", NovelID)).Add("expand", "originNeedFireMoney").NewRequests().
		Unmarshal(&Catalogue)

	return Catalogue

}

func GET_CHAPTER_CONTENT(chapter_id string) BoluobaoStructs.Content {
	var Content BoluobaoStructs.Content
	request.Get("Chaps/"+chapter_id).Add("expand", "content").NewRequests().Unmarshal(&Content)
	return Content
}

func GET_SEARCH(keyword string, page int) BoluobaoStructs.Search {
	var Search BoluobaoStructs.Search
	request.Get("search/novels/result").Add("q", keyword).
		Add("size", "20").Add("page", strconv.Itoa(page)).NewRequests().Unmarshal(&Search)
	return Search
}

func LOGIN_ACCOUNT(username, password string) string {
	CookieJar := request.Post("sessions").Add("username", username).
		Add("password", password).NewRequests().Unmarshal(&BoluobaoStructs.Login).GetCookie()
	for _, cookie := range CookieJar {
		BoluobaoStructs.Login.Cookie += cookie.Name + "=" + cookie.Value + ";"
	}

	if BoluobaoStructs.Login.Status.HTTPCode == 200 {
		return BoluobaoStructs.Login.Cookie
	} else {
		if message := BoluobaoStructs.Login.Status.Msg.(string); message == "用户名密码不匹配" {
			fmt.Println(message)
		} else {
			fmt.Println("login failed! error:", message)
		}
	}
	return ""
}
