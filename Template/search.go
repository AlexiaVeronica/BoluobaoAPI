package Template

type Search struct {
	Status status `json:"status"`
	Data   struct {
		Novels []BookInfoData `json:"novels"`
	} `json:"data"`
}
