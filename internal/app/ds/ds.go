package ds

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/staurran/messengerKR.git/internal/app/constProject"
)

type Chat struct {
	ChatId      uint                  `sql:"type:uuid;primary_key;default:" json:"chatId" gorm:"primarykey"`
	Type        constProject.TypeChat `json:"type"`
	ChatName    string                `json:"chatName"`
	Avatar      string                `json:"avatar"`
	Description string                `json:"description"`
	Image       string                `json:"image"`
}

type Users struct {
	UserId      uint              `sql:"type:uuid;primary_key;default:" json:"userId" gorm:"primarykey"`
	Role        constProject.Role `json:"role"`
	Username    string            `json:"username"`
	Password    string            `json:"password"`
	Avatar      string            `json:"avatar"`
	PhoneNumber string            `json:"phoneNumber"`
	Bio         string            `json:"bio"`
}

type ChatUser struct {
	ChatUserId uint                  `sql:"type:uuid;primary_key;default:" json:"chatUserId" gorm:"primarykey"`
	UserId     uint                  `json:"userId"`
	ChatId     uint                  `json:"chatId"`
	ChatRole   constProject.ChatRole `json:"chatRole"`
}

type JWTClaims struct {
	jwt.StandardClaims
	UserId uint `json:"userId"`
	Role   constProject.Role
}

type Message struct {
	MessageId      uint      `sql:"type:uuid;primary_key;default:" json:"messageId" gorm:"primarykey"`
	Context        string    `json:"context"`
	UserFrom       uint      `json:"userFrom"`
	UserTo         uint      `json:"userTo"`
	OriginUserFrom int       `json:"originUserFrom"`
	TimeCreated    time.Time `json:"timeCreated"`
	TimeUpdated    time.Time `json:"timeUpdated"`
}

type PhotoMess struct {
	PhotoMessId uint   `sql:"type:uuid;primary_key;default:" json:"photoMessId" gorm:"primarykey"`
	MessageId   uint   `json:"messageId"`
	Link        string `json:"link"`
}

type AudioMess struct {
	AudioMessId uint   `sql:"type:uuid;primary_key;default:" json:"audioMessId" gorm:"primarykey"`
	MessageId   uint   `json:"messageId"`
	Link        string `json:"link"`
}

type GoodOrder struct {
	Id_row   uint `sql:"type:uuid;primary_key;default:" json:"Id_row" gorm:"primarykey"`
	Id_good  uint `json:"id_good"`
	Quantity int  `json:"quantity"`
	Id_order uint `json:"id_order"`
}

type Statuses struct {
	Id_status   uint   `sql:"type:uuid;primary_key;default:" json:"Id_status" gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
