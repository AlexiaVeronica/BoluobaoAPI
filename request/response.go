package request

import "github.com/VeronicaAlexia/BoluobaoAPI/config"

func ApiHost() string {
	var Host string
	if config.AppConfig.App {
		Host = "https://api.sfacg.com/"
	} else {
		Host = "https://minipapi.sfacg.com/pas/mpapi/"
	}
	return Host
}

func Get(path string) *HttpUtils {
	return NewHttpUtils(ApiHost(), path, "GET")
}
func Post(path string) *HttpUtils {
	return NewHttpUtils(ApiHost(), path, "POST")
}
func Put(path string) *HttpUtils {
	return NewHttpUtils(ApiHost(), path, "PUT")
}
