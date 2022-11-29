package Template

type FXrecommend struct {
	Status Status `json:"status"`
	Data   struct {
		HotPush []struct {
			Novel struct {
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
					SysTags []struct {
						SysTagID int    `json:"sysTagId"`
						TagName  string `json:"tagName"`
					} `json:"sysTags"`
					Discount           float64 `json:"discount"`
					DiscountExpireDate string  `json:"discountExpireDate"`
					HomeFlag           []struct {
						Desc      string `json:"desc"`
						Type      int    `json:"type"`
						DateRange int    `json:"dateRange"`
						Num       int    `json:"num"`
					} `json:"homeFlag"`
				} `json:"expand"`
			} `json:"novel"`
			BeginDate string `json:"beginDate"`
		} `json:"hotPush"`
	} `json:"data"`
}
type NewBookRecommend struct {
	Status Status `json:"status"`
	Data   struct {
		NewPush []struct {
			Novel struct {
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
					SysTags []struct {
						SysTagID int    `json:"sysTagId"`
						TagName  string `json:"tagName"`
					} `json:"sysTags"`
					Discount           float64 `json:"discount"`
					DiscountExpireDate string  `json:"discountExpireDate"`
					HomeFlag           []struct {
						Desc      string `json:"desc"`
						Type      int    `json:"type"`
						DateRange int    `json:"dateRange"`
						Num       int    `json:"num"`
					} `json:"homeFlag"`
				} `json:"expand"`
			} `json:"novel"`
			BeginDate string `json:"beginDate"`
		} `json:"newPush"`
	} `json:"data"`
}
type Actpush struct {
	Status Status `json:"status"`
	Data   []struct {
		ImgURL  string `json:"imgUrl"`
		Link    string `json:"link"`
		Type    string `json:"type"`
		Desc    string `json:"desc"`
		EndDate string `json:"endDate"`
	} `json:"data"`
}
