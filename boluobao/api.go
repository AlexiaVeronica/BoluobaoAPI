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

func GET_BOOK_SHELF_INFORMATION() BoluobaoStructs.InfoData {
	var bookshelfData BoluobaoStructs.InfoData
	request.Get("user/Pockets").Add("expand", "novels").NewRequests().Unmarshal(&bookshelfData)
	return bookshelfData
}

func GET_SYS_TAG_LIST() BoluobaoStructs.SysTags {
	var SysTags BoluobaoStructs.SysTags
	request.Get("novels/0/sysTags").Add("filter", "push").NewRequests().Unmarshal(&SysTags)
	return SysTags
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
