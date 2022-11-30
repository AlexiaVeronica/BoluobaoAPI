package Template

type SysTags struct {
	Status Status `json:"status"`
	Data   []struct {
		SysTagID    int    `json:"sysTagId"`
		RefferTimes int    `json:"refferTimes"`
		TagName     string `json:"tagName"`
		ImageURL    string `json:"imageUrl"`
		NovelCount  int    `json:"novelCount"`
		IsDefault   bool   `json:"isDefault"`
	} `json:"data"`
}

type Tag struct {
	Status Status `json:"status"`
	Data   []struct {
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
			TypeName string `json:"typeName"`
			SysTags  []struct {
				SysTagID int    `json:"sysTagId"`
				TagName  string `json:"tagName"`
			} `json:"sysTags"`
			Discount           float64 `json:"discount"`
			DiscountExpireDate string  `json:"discountExpireDate"`
		} `json:"expand"`
	} `json:"data"`
}
