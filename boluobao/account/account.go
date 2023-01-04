package account

import "fmt"

type User struct{}

func (u *User) UserInformation() {
	response := GET_ACCOUNT_INFORMATION()
	if response.Status.HTTPCode == 200 {
		fmt.Println("用户名:", response.Data.NickName)
		//response2 := GET_ACCOUNT_MONEY_INFORMATION()
		//fmt.Println("余额:", response2.Data.)

	} else {
		fmt.Println("获取用户信息失败:", response.Status.Msg)
	}
}
