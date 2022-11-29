package request

type AppRequest struct{ App bool }

var ApiHost string

var Cookie string

func (a *AppRequest) SetApiHost() string {
	if a.App {
		ApiHost = "https://api.sfacg.com/"
	} else {
		ApiHost = "https://minipapi.sfacg.com/pas/mpapi/"
	}
	return ApiHost
}

func Get(path string) *HttpUtils {
	return NewHttpUtils(ApiHost, path, "GET")
}
func Post(path string) *HttpUtils {
	return NewHttpUtils(ApiHost, path, "POST")
}
func Put(path string) *HttpUtils {
	return NewHttpUtils(ApiHost, path, "PUT")
}
