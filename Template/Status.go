package Template

type Status struct {
	HTTPCode  int         `json:"httpCode"`
	ErrorCode int         `json:"errorCode"`
	MsgType   int         `json:"msgType"`
	Msg       interface{} `json:"msg"`
}
