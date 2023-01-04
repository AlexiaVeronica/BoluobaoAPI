package account

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
)

type User struct{}

func (u *User) UserInformation() *Template.AccountData {
	response := GET_ACCOUNT_INFORMATION()
	if response.Status.HTTPCode == 200 {
		return &response.Data
	} else {
		fmt.Println("获取用户信息失败:", response.Status.Msg)
		return nil
	}
}

func (u *User) ShowUserInformation(AccountData *Template.AccountData) {
	fmt.Println("用户名:", AccountData.NickName)
	fmt.Println("用户ID:", AccountData.AccountID)
	//fmt.Println("用户等级:", AccountData)

}

func (u *User) UserLogin(username string, password string, retry int) {
	if retry <= 0 {
		retry = 1
	} else if retry > 3 {
		retry = 3
	}
	for i := 0; i < retry; i++ {
		response := LOGIN_ACCOUNT(username, password)
		if response != "" {
			fmt.Println("登录成功")
			break
		} else {
			fmt.Println("retry login account:", i+1)
		}
	}
}
