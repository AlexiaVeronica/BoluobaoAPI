package BoluobaoStructs

var Login = struct {
	Status status `json:"status"`
	Cookie string
	Data   interface{} `json:"data"`
}{}
