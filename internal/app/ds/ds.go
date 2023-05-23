package ds

import (
	"time"

	"github.com/staurran/messengerKR.git/internal/app/constProject"
)

type Chat struct {
	Id          uint   `sql:"type:uuid;primary_key;default:" json:"chatId" gorm:"primarykey"`
	Name        string `json:"chatName"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}

type User struct {
	Id       uint   `sql:"type:uuid;primary_key;default:" json:"userId" gorm:"primarykey" structs:"id"`
	Username string `json:"username" structs:"username"`
	Password string `json:"password" structs:"password"`
	Avatar   string `json:"avatar" structs:"avatar"`
	Phone    string `json:"phoneNumber" structs:"phoneNumber"`
	Bio      string `json:"bio" structs:"bio"`
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
}

type Shown struct {
	Id        uint `sql:"type:uuid;primary_key;default:" json:"shownId" gorm:"primarykey"`
	MessageId uint
	UserId    uint
	Shown     bool
}
