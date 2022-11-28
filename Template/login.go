package Template

var Login = struct {
	Status status `json:"status"`
	Cookie string
	Data   interface{} `json:"data"`
}{}
