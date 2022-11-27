package BoluobaoStructs

type SysTags struct {
	Status status `json:"status"`
	Data   []struct {
		SysTagID    int    `json:"sysTagId"`
		RefferTimes int    `json:"refferTimes"`
		TagName     string `json:"tagName"`
		ImageURL    string `json:"imageUrl"`
		NovelCount  int    `json:"novelCount"`
		IsDefault   bool   `json:"isDefault"`
	} `json:"data"`
}
