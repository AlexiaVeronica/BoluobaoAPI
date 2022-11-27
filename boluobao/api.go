package boluobao

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/BoluobaoStructs"
	"github.com/VeronicaAlexia/BoluobaoAPI/BoluobaoStructs/bookshelf"
	request2 "github.com/VeronicaAlexia/BoluobaoAPI/boluobao/request"
	url_ "net/url"
	"os"
	"strconv"
)

func GET_BOOK_INFORMATION(NovelId string) BoluobaoStructs.BookInfo {
	var BookInfo BoluobaoStructs.BookInfo
	request2.Get("novels/"+NovelId).Add("expand", "intro,tags,sysTags").NewRequests().
		Unmarshal(&BookInfo)
	return BookInfo
}

func GET_ACCOUNT_INFORMATION() BoluobaoStructs.Account {
	var Account BoluobaoStructs.Account
	request2.Get("user").NewRequests().Unmarshal(&Account)
	return Account
}

func GET_BOOK_SHELF_INFORMATION() bookshelf.InfoData {
	var bookshelfData bookshelf.InfoData
	request2.Get("user/Pockets").Add("expand", "novels").NewRequests().Unmarshal(&bookshelfData)
	return bookshelfData
}

func GET_CATALOGUE(NovelID string) BoluobaoStructs.Catalogue {
	var Catalogue BoluobaoStructs.Catalogue
	request2.Get(fmt.Sprintf("novels/%v/dirs", NovelID)).Add("expand", "originNeedFireMoney").NewRequests().
		Unmarshal(&Catalogue)

	return Catalogue

}

func GET_CHAPTER_CONTENT(chapter_id string) BoluobaoStructs.Content {
	var Content BoluobaoStructs.Content
	request2.Get("Chaps/"+chapter_id).Add("expand", "content").NewRequests().Unmarshal(&Content)
	return Content
}

func GET_SEARCH(keyword string, page int) BoluobaoStructs.Search {
	var Search BoluobaoStructs.Search
	request2.Get("search/novels/result").Add("q", url_.QueryEscape(keyword)).
		Add("size", "20").Add("page", strconv.Itoa(page)).NewRequests().Unmarshal(&Search)
	return Search

}

func LOGIN_ACCOUNT(username, password string) string {
	CookieJar := request2.NewHttpUtils("sessions", "POST").Add("username", username).
		Add("password", password).NewRequests().Unmarshal(&BoluobaoStructs.Login).GetCookie()
	for _, cookie := range CookieJar {
		BoluobaoStructs.Login.Cookie += cookie.Name + "=" + cookie.Value + ";"
	}

	if BoluobaoStructs.Login.Status.HTTPCode == 200 {
		return BoluobaoStructs.Login.Cookie
	} else {
		if message := BoluobaoStructs.Login.Status.Msg.(string); message == "用户名密码不匹配" {
			fmt.Println(message)
			os.Exit(0)
		} else {
			fmt.Println("login failed! error:", message)
		}
	}
	return ""
}

func GET_TASK_LIST() string {
	s := new(BoluobaoStructs.Task)
	request2.Get("user/tasks/").Add(
		"taskCategory", "1").Add(
		"package", "com.sfacg").Add(
		"deviceToken", "").Add(
		"page", "0").Add(
		"size", "20").NewRequests().Unmarshal(s)
	if s != nil && s.Status.HttpCode == 200 {
		for _, task_info := range s.Data {
			fmt.Println("任务号:", task_info.TaskId, "\t\t\t任务名:", task_info.Name)
		}
	} else {
		fmt.Println("task error:", s.Status.Msg)
	}
	return ""
}
