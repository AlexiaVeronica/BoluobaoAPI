package tag

import "strconv"

func TagsParams(TagID string, page int, size string) map[string]string {
	return map[string]string{
		"systagids":      TagID,
		"isfree":         "both",
		"size":           size,
		"charcountbegin": "0",
		"updatedays":     "-1",
		"expand":         "chapterCount,bigBgBanner,bigNovelCover,typeName,intro,fav,ticket,pointCount,tags,sysTags,signlevel,rankinglist,firstchapter,latestchapter,essaytag,auditCover,preOrderInfo,customTag,topic,isbranch",
		"sort":           "latest",
		"page":           strconv.Itoa(page),
		"charcountend":   "0",
		"isfinish":       "both",
	}
}
