package user

type Contact struct {
	ContactId uint   `json:"contactId"`
	UserId    uint   `json:"userId"`
	Avatar    string `json:"avatar"`
	Username  string `json:"username"`
	Phone     string `json:"phone"`
}
