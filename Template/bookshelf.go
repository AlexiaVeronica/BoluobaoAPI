package Template

type ShelfData struct {
	AccountID  int    `json:"accountId"`
	PocketID   int    `json:"pocketId"`
	Name       string `json:"name"`
	TypeID     int    `json:"typeId"`
	CreateTime string `json:"createTime"`
	IsFull     bool   `json:"isFull"`
	CanModify  bool   `json:"canModify"`
	Expand     struct {
		Novels []BookInfoData `json:"novels"`
	} `json:"expand"`
}

type InfoData struct {
	Status Status      `json:"status"`
	Data   []ShelfData `json:"data"`
}
