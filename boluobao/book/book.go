package book

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
)

func GET_BOOK_INFORMATION(NovelId string) Template.BookInfo {
	var BookInfo Template.BookInfo
	expand := "chapterCount,bigBgBanner,bigNovelCover,typeName,intro,fav,ticket,pointCount,tags,sysTags,signlevel,discount,discountExpireDate,totalNeedFireMoney,rankinglist,originTotalNeedFireMoney,firstchapter,latestchapter,latestcommentdate,essaytag,auditCover,preOrderInfo,customTag,topic,unauditedCustomtag,homeFlag,isbranch"
	request.Get("novels/"+NovelId).Add("expand", expand).NewRequests().Unmarshal(&BookInfo)
	return BookInfo
}

func GET_CATALOGUE(NovelID string) Template.Catalogue {
	var Catalogue Template.Catalogue
	request.Get(fmt.Sprintf("novels/%v/dirs", NovelID)).Add("expand", "originNeedFireMoney").NewRequests().
		Unmarshal(&Catalogue)
	return Catalogue

}

func GET_CHAPTER_CONTENT(chapter_id string) Template.Content {
	var Content Template.Content
	request.Get("Chaps/"+chapter_id).Add("expand", "content").NewRequests().Unmarshal(&Content)
	return Content
}
