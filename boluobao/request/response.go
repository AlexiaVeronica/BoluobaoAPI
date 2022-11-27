package request

func Get(path string) *HttpUtils {
	return NewHttpUtils(APP_TYPE.Host, path, "GET")
}
func Post(path string) *HttpUtils {
	return NewHttpUtils(APP_TYPE.Host, path, "POST")
}
func Put(path string) *HttpUtils {
	return NewHttpUtils(APP_TYPE.Host, path, "PUT")
}
