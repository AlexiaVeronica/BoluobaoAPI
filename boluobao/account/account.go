package account

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"strconv"
)

func GET_ACCOUNT_INFORMATION() Template.Account {
	var Account Template.Account
	request.Get("user").NewRequests().Unmarshal(&Account)
	return Account
}

func GET_ACCOUNT_MONEY_INFORMATION() Template.Account {
	var Account Template.Account
	request.Get("user/money").NewRequests().Unmarshal(&Account)
	return Account
}

func GET_ACCOUNT_IP() Template.AccountIP {
	var AccountIP Template.AccountIP
	request.Get("position").NewRequests().Unmarshal(&AccountIP)
	return AccountIP
}

// GET_USER_INFORMATION 公开用户信息
func GET_USER_INFORMATION(AccountId string) Template.Users {
	var Users Template.Users
	params := map[string]string{"expand": "introduction,bigAvatar,avatar,backgroundPic,fansNum,followNum,followyou,youfollow,verifyType,verifyInfo,avatarFrame,youblock,widgets,growup"}
	request.Get("users/" + AccountId).NewRequests().AddAll(params).Unmarshal(&Users)
	return Users
}

// GET_USER_WORKS 用户作品信息
func GET_USER_WORKS(AccountId string) Template.AuthorInfo {
	var AuthorInfo Template.AuthorInfo
	params := map[string]string{"expand": "typeName,sysTags,isbranch"}
	request.Get("users/" + AccountId + "/novels").NewRequests().AddAll(params).Unmarshal(&AuthorInfo)
	return AuthorInfo
}

// GET_USER_COMMENT 用户评论信息
func GET_USER_COMMENT(AccountId string, page int) Template.AccountComment {
	var AccountComment Template.AccountComment
	params := map[string]string{"expand": "novels,comics,albums,tags,sysTags,authorName", "size": "20", "page": strconv.Itoa(page)}
	request.Get("users/" + AccountId + "/dynamics").NewRequests().AddAll(params).Unmarshal(&AccountComment)
	return AccountComment
}

func LOGIN_ACCOUNT(username, password string) string {
	var Cookie string
	params := map[string]string{"username": username, "password": password}
	CookieJar := request.Post("sessions").AddAll(params).NewRequests().Unmarshal(&Template.Login).GetCookie()
	if Template.Login.Status.HTTPCode == 200 {
		for _, cookie := range CookieJar {
			Cookie += cookie.Name + "=" + cookie.Value + ";"
		}
		return Cookie
	} else {
		fmt.Println("login failed! error:", Template.Login.Status.Msg)
		return ""
	}
}
