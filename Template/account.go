package Template

type Users struct {
	Status Status `json:"status"`
	Data   struct {
		AccountID int    `json:"accountId"`
		UserName  string `json:"userName"`
		NickName  string `json:"nickName"`
	} `json:"data"`
}
type Account struct {
	Status Status `json:"status"`
	Data   struct {
		UserName     string `json:"userName"`
		NickName     string `json:"nickName"`
		Email        string `json:"email"`
		AccountID    int    `json:"accountId"`
		RoleName     string `json:"roleName"`
		FireCoin     int    `json:"fireCoin"`
		Avatar       string `json:"avatar"`
		IsAuthor     bool   `json:"isAuthor"`
		PhoneNum     string `json:"phoneNum"`
		RegisterDate string `json:"registerDate"`
	} `json:"data"`
}
type AuthorInfo struct {
	Status Status         `json:"status"`
	Data   []BookInfoData `json:"data"`
}
