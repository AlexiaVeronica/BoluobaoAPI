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
	Status Status         `json:"status"`
	Data   []BookInfoData `json:"data"`
}

type SysTagList struct {
	SysTagID int    `json:"sysTagId"`
	TagName  string `json:"tagName"`
}
