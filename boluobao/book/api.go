package book

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
)

func GET_BOOK_INFORMATION(NovelId string) Template.BookInfo {
	var BookInfoType Template.BookInfo
	expand := "chapterCount,bigBgBanner,bigNovelCover,typeName,intro,fav,ticket,pointCount,tags,sysTags,signlevel,discount,discountExpireDate,totalNeedFireMoney,rankinglist,originTotalNeedFireMoney,firstchapter,latestchapter,latestcommentdate,essaytag,auditCover,preOrderInfo,customTag,topic,unauditedCustomtag,homeFlag,isbranch"
	request.Get("novels/"+NovelId).Add("expand", expand).NewRequests().Unmarshal(&BookInfoType)
	return BookInfoType
}

func GET_CATALOGUE(NovelID string) Template.Catalogue {
	var CatalogueType Template.Catalogue
	request.Get(fmt.Sprintf("novels/%v/dirs", NovelID)).Add("expand", "originNeedFireMoney").NewRequests().
		Unmarshal(&CatalogueType)
	return CatalogueType

}

func GET_CHAPTER_CONTENT(chapter_id string) Template.Content {
	var ContentType Template.Content
	request.Get("Chaps/"+chapter_id).Add("expand", "content").NewRequests().Unmarshal(&ContentType)
	return ContentType
}
