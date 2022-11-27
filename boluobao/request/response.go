package request

func Get(api_url string) *HttpUtils {
	return NewHttpUtils(api_url, "GET")
}
func Post(api_url string) *HttpUtils {
	return NewHttpUtils(api_url, "POST")
}
func Put(api_url string) *HttpUtils {
	return NewHttpUtils(api_url, "PUT")
}
