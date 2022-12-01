package Template

type FXrecommend struct {
	Status Status `json:"status"`
	Data   struct {
		HotPush []struct {
			Novel     BookInfoData `json:"novel"`
			BeginDate string       `json:"beginDate"`
		} `json:"hotPush"`
	} `json:"data"`
}
type NewBookRecommend struct {
	Status Status `json:"status"`
	Data   struct {
		NewPush []struct {
			Novel     BookInfoData `json:"novel"`
			BeginDate string       `json:"beginDate"`
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
