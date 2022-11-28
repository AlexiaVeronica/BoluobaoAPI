package Template

type Rank struct {
	Status struct {
		HTTPCode  int         `json:"httpCode"`
		ErrorCode int         `json:"errorCode"`
		MsgType   int         `json:"msgType"`
		Msg       interface{} `json:"msg"`
	} `json:"status"`
	Data []struct {
		AuthorID       int     `json:"authorId"`
		LastUpdateTime string  `json:"lastUpdateTime"`
		MarkCount      int     `json:"markCount"`
		NovelCover     string  `json:"novelCover"`
		BgBanner       string  `json:"bgBanner"`
		NovelID        int     `json:"novelId"`
		NovelName      string  `json:"novelName"`
		Point          float64 `json:"point"`
		IsFinish       bool    `json:"isFinish"`
		AuthorName     string  `json:"authorName"`
		CharCount      int     `json:"charCount"`
		ViewTimes      int     `json:"viewTimes"`
		TypeID         int     `json:"typeId"`
		AllowDown      bool    `json:"allowDown"`
		AddTime        string  `json:"addTime"`
		IsSensitive    bool    `json:"isSensitive"`
		SignStatus     string  `json:"signStatus"`
		CategoryID     int     `json:"categoryId"`
		Expand         struct {
			TypeName      string        `json:"typeName"`
			Tags          []interface{} `json:"tags"`
			Ticket        int           `json:"ticket"`
			LatestChapter struct {
				Title   string `json:"title"`
				ChapID  int    `json:"chapId"`
				AddTime string `json:"addTime"`
			} `json:"latestChapter"`
			Bonus    int     `json:"bonus"`
			Discount float64 `json:"discount"`
		} `json:"expand"`
	} `json:"data"`
}
