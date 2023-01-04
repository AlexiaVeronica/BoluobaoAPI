package request

var AppConfig = Config{}

type Config struct {
	App      bool
	Cookie   string
	AppKey   string
	DeviceId string
}

func ApiHost() string {
	var Host string
	if AppConfig.App {
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
