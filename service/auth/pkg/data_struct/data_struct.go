package dataStruct

type User struct {
	Id       uint   `sql:"unique;type:uuid;primary_key;servicedefault:" json:"userId" gorm:"primaryKey;unique"`
	Phone    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	BirthDay string `json:"birthDay" sql:"type:date" gorm:"type:date"`
	Bio      string `json:"bio"`
}
