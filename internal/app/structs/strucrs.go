package structs

import "github.com/staurran/messengerKR.git/internal/app/constProject"

type ChatStruct struct {
	Id                  uint `json:"id"`
	Name                string
	Avatar              string
	QuantityUnshownMess int64 `json:"quantityUnshownMess"`
	LastMessage
}

type LastMessage struct {
	Username string
	Content  string
}

type ChatRepoStruct struct {
	Id       uint `json:"id"`
	Name     string
	Avatar   string
	CountMes int64 `json:"count_mes"`
}

type GetChatsResult struct {
	chats []ChatStruct
}

type InpCreateChat struct {
	Users       []UserChat            `json:"users"`
	Name        string                `json:"name"`
	Avatar      string                `json:"avatar"`
	Description string                `json:"description"`
	Type        constProject.TypeChat `json:"type"`
}

type UserChat struct {
	UserId   uint                  `json:"userId"`
	Username string                `json:"username"`
	Role     constProject.ChatRole `json:"role"`
}

type Message struct {
	ID           uint   `json:"id"`
	Content      string `json:"content"`
	UserFrom     string `json:"userFrom"`
	OriginalUser string `json:"originalUser"`
	Attachment   []string
	Photos       []string
	Audio        string
	TimeCreated  string
	Updated      bool
	Reactions    []ReactionMes
	Shown        bool
}

type ReactionMes struct {
	Reaction string
	Quantity int
}
