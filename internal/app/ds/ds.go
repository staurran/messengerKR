package ds

import (
	"time"

	"github.com/staurran/messengerKR.git/internal/app/constProject"
)

type Chat struct {
	Id          uint   `sql:"type:uuid;primary_key;default:" json:"chatId" gorm:"primarykey"`
	Name        string `json:"chatName"`
	Avatar      uint   `json:"avatar"`
	Description string `json:"description"`
}

type User struct {
	Id       uint   `sql:"type:uuid;primary_key;default:" json:"userId" gorm:"primarykey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phoneNumber"`
	Bio      string `json:"bio"`
}

type UserPhoto struct {
	Id     uint `sql:"unique;type:uuid;primary_key;servicedefault:" json:"userPhotoId" gorm:"primaryKey;unique"`
	UserId uint `json:"userId" gorm:"foreignKey"`
	Photo  uint `json:"photo"`
	Avatar bool `json:"avatar"`
}

type Contact struct {
	Id        uint `sql:"type:uuid;primary_key;default:" json:"userId" gorm:"primarykey"`
	UserID    uint `json:"userID"`
	ContactID uint `json:"contactID"`
}

type ChatUser struct {
	Id       uint                  `sql:"type:uuid;primary_key;default:" json:"chatUserId" gorm:"primarykey"`
	UserID   uint                  `json:"userId"`
	ChatID   uint                  `json:"chatId"`
	ChatRole constProject.ChatRole `json:"chatRole"`
}

type Photo struct {
	Id        uint   `sql:"type:uuid;primary_key;default:" json:"photoMessId" gorm:"primarykey"`
	MessageID uint   `json:"messageId"`
	Photo     string `json:"photo"`
}

type Audio struct {
	Id        uint   `sql:"type:uuid;primary_key;default:" json:"audioMessId" gorm:"primarykey"`
	MessageID uint   `json:"messageId"`
	Audio     string `json:"audio"`
}

type Attachment struct {
	Id         uint   `sql:"type:uuid;primary_key;default:" json:"attachmentMessId" gorm:"primarykey"`
	MessageID  uint   `json:"messId"`
	Attachment string `json:"attachment"`
}

type Reaction struct {
	Id        uint   `sql:"type:uuid;primary_key;default:" json:"reactionId" gorm:"primarykey"`
	MessageID uint   `json:"messId"`
	Name      string `json:"reactionName"`
	UserID    uint   `json:"userId"`
}

type Message struct {
	Id      uint `sql:"type:uuid;primary_key;default:" json:"messageId" gorm:"primarykey"`
	Content string
	Time    time.Time
	UserId  uint
	ChatId  uint
	Shown   bool
}

type Shown struct {
	Id        uint `sql:"type:uuid;primary_key;default:" json:"shownId" gorm:"primarykey"`
	MessageId uint
	UserId    uint
	Shown     bool
}
