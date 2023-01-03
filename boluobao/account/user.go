package account

import "fmt"

func UserInformation() {
	response := GET_ACCOUNT_INFORMATION()
	if response.Status.HTTPCode == 200 {
		fmt.Println("用户信息:", response.Data.NickName)

	} else {
		fmt.Println("获取用户信息失败:", response.Status.Msg)
	}
}
