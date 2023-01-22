package api

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
)

var Request = Template.ConfigRequest{}

func Host() string {
	var host string
	if Request.App {
		host = "https://api.sfacg.com/"
	} else {
		host = "https://minipapi.sfacg.com/pas/mpapi/"
	}
	return host
}

func Get(path string) *request.HttpUtils {
	return request.NewHttpUtils(Host(), path, "GET", Request)
}
func Post(path string) *request.HttpUtils {
	return request.NewHttpUtils(Host(), path, "POST", Request)
}
func Put(path string) *request.HttpUtils {
	return request.NewHttpUtils(Host(), path, "PUT", Request)
}
