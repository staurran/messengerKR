package user

type Contact struct {
	ContactId uint
	UserId    uint
	Avatar    string
	Name      string
	Phone     string
}

type UserInfo struct {
	UserId uint
	Name   string
	Phone  string
	Photo  []string
	Bio    string
}
