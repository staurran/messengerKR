package dataStruct

type User struct {
	Id       uint   `sql:"unique;type:uuid;primary_key;servicedefault:" json:"userId" gorm:"primaryKey;unique"`
	Phone    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Avatar   string `json:"avatar"`
}
