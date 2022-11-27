package BoluobaoStructs

type status struct {
	HTTPCode  int         `json:"httpCode"`
	ErrorCode int         `json:"errorCode"`
	MsgType   int         `json:"msgType"`
	Msg       interface{} `json:"msg"`
}
