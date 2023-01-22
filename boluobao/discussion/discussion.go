package discussion

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/api"
	"strconv"
)

// Feeds 获取讨论区动态
func Feeds(page int, filter string) *Template.Discussion {
	// filter:[latest, hot]
	if filter == "" {
		filter = "hot"
	}
	var Discussion Template.Discussion
	params := map[string]string{"filter": filter, "expand": "novels,comics,albums,tags,sysTags,authorName,hasfollowed", "page": strconv.Itoa(page), "size": "20"}
	api.Get("user/feeds").AddAll(params).NewRequests().Unmarshal(&Discussion)
	return &Discussion
}
