package request

func Get(path string) *HttpUtils {
	return NewHttpUtils("https://minipapi.sfacg.com/pas/mpapi/", path, "GET")
}
func Post(path string) *HttpUtils {
	return NewHttpUtils("https://minipapi.sfacg.com/pas/mpapi/", path, "POST")
}
func Put(path string) *HttpUtils {
	return NewHttpUtils("https://minipapi.sfacg.com/pas/mpapi/", path, "PUT")
}
