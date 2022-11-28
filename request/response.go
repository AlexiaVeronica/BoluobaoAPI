package request

type AppRequest struct{ App bool }

var api_host string

func (a *AppRequest) SetApiHost() {
	if a.App {
		api_host = "https://api.sfacg.com/"
	} else {
		api_host = "https://minipapi.sfacg.com/pas/mpapi/"
	}
}

func Get(path string) *HttpUtils {
	return NewHttpUtils(api_host, path, "GET")
}
func Post(path string) *HttpUtils {
	return NewHttpUtils(api_host, path, "POST")
}
func Put(path string) *HttpUtils {
	return NewHttpUtils(api_host, path, "PUT")
}
