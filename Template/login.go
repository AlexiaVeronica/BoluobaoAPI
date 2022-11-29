package Template

var Login = struct {
	Status Status `json:"status"`
	Cookie string
	Data   interface{} `json:"data"`
}{}
