package chat

import "github.com/staurran/messengerKR.git/internal/app/ds"

type IRepositoryPhoto interface {
	CreateChat(chat *ds.Chat) error
	SaveChatUsers(user []ds.ChatUser) error
	DeleteChatUser(user ds.ChatUser) error
	CheckAdmin(userId, chatId uint) error
	ChangeChat(chat ds.Chat) error
}
