package Template

type Rank struct {
	Status Status         `json:"status"`
	Data   []BookInfoData `json:"data"`
}
