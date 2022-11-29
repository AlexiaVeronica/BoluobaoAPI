package Template

type Search struct {
	Status Status `json:"status"`
	Data   struct {
		Novels []BookInfoData `json:"novels"`
	} `json:"data"`
}
