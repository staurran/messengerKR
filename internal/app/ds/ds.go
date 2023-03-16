package ds

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/staurran/messengerKR.git/internal/app/constProject"
)

type Chat struct {
	ID          uint                  `sql:"type:uuid;primary_key;default:" json:"chatId" gorm:"primarykey"`
	Type        constProject.TypeChat `json:"type"`
	Name        string                `json:"chatName"`
	Avatar      string                `json:"avatar"`
	Description string                `json:"description"`
}

type Users struct {
	ID       uint              `sql:"type:uuid;primary_key;default:" json:"userId" gorm:"primarykey"`
	Role     constProject.Role `json:"role"`
	Username string            `json:"username"`
	Password string            `json:"password"`
	Avatar   string            `json:"avatar"`
	Phone    string            `json:"phoneNumber"`
	Bio      string            `json:"bio"`
}

type ChatUser struct {
	ID       uint                  `sql:"type:uuid;primary_key;default:" json:"chatUserId" gorm:"primarykey"`
	UserID   uint                  `json:"userId"`
	ChatID   uint                  `json:"chatId"`
	ChatRole constProject.ChatRole `json:"chatRole"`
}

type JWTClaims struct {
	jwt.StandardClaims
	UserID uint `json:"userId"`
	Role   constProject.Role
}

type Message struct {
	ID           uint      `sql:"type:uuid;primary_key;default:" json:"messageId" gorm:"primarykey"`
	Context      string    `json:"context"`
	UserFromID   uint      `json:"userFrom"`
	UserToID     uint      `json:"userTo"`
	OriginUserID uint      `json:"originUserFrom"`
	TimeCreated  time.Time `json:"timeCreated"`
	TimeUpdated  time.Time `json:"timeUpdated"`
}

type PhotoMess struct {
	ID        uint   `sql:"type:uuid;primary_key;default:" json:"photoMessId" gorm:"primarykey"`
	MessageID uint   `json:"messageId"`
	Link      string `json:"link"`
}

type AudioMess struct {
	ID        uint   `sql:"type:uuid;primary_key;default:" json:"audioMessId" gorm:"primarykey"`
	MessageID uint   `json:"messageId"`
	Link      string `json:"link"`
}

type AttachmentMess struct {
	ID        uint   `sql:"type:uuid;primary_key;default:" json:"attachmentMessId" gorm:"primarykey"`
	MessageID uint   `json:"messId"`
	Link      string `json:"link"`
}

type MessageShown struct {
	ID        uint   `sql:"type:uuid;primary_key;default:" json:"messShowId" gorm:"primarykey"`
	MessageID uint   `json:"messId"`
	UserID    uint   `json:"userId"`
	Shown     string `json:"shown"`
}

type Reaction struct {
	ID        uint   `sql:"type:uuid;primary_key;default:" json:"reactionId" gorm:"primarykey"`
	MessageID uint   `json:"messId"`
	Name      string `json:"reactionName"`
	UserID    string `json:"userId"`
}
