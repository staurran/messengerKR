package app

import (
	"github.com/staurran/messengerKR.git/internal/app/constProject"
	"github.com/staurran/messengerKR.git/internal/app/repository"
)

////Users

type CurrentUserResult struct {
	Username string
	Avatar   string
	Phone    string
	Bio      string
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone" binding:"required"`
	Bio      string `json:"bio"`
}

type ContactUser struct {
	Phone string
}

///Chats

type ChatStruct struct {
	Id                  uint `json:"id"`
	Name                string
	Avatar              string
	QuantityUnshownMess int64 `json:"quantityUnshownMess"`
	repository.LastMessage
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
