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
