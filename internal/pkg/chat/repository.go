package chat

import "github.com/staurran/messengerKR.git/internal/app/ds"

type IRepositoryPhoto interface {
	CreateChat(chat *ds.Chat) error
	SaveChatUsers(user []ds.ChatUser) error
	DeleteChatUser(user ds.ChatUser) error
	CheckAdmin(userId, chatId uint) error
	ChangeChat(chat ds.Chat) error
	GetChats(userId uint) ([]ChatRepoStruct, error)
	GetLastMes(chatId uint) (LastMessage, error)
	CreateMessage(message *ds.Message) error
	CreateMesUserShown(msg []ds.Shown) error
	GetChatUsers(chatId uint) ([]uint, error)
	SaveAttachments([]ds.Attachment) error
	SavePhotos(photos []ds.Photo) error
	SaveAudio(audio ds.Audio) error
}
