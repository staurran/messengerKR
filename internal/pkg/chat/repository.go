package chat

import "github.com/staurran/messengerKR.git/internal/app/ds"

type IRepositoryChat interface {
	CreateChat(chat *ds.Chat) error
	SaveChatUsers(user []ds.ChatUser) error
	DeleteChatUser(user *ds.ChatUser) error
	CheckAdmin(userId, chatId uint) error
	ChangeChat(chat ds.Chat) error
	GetChats(userId uint) ([]ChatRepoStruct, error)
	GetLastMes(chatId uint) (LastMessage, error)
	CreateMessage(message *ds.Message) error
	CreateMesUserShown(msg []ds.Shown) error
	GetChatUsers(chatId uint) ([]uint, error)
	SaveAttachments([]ds.Attachment) error
	SavePhoto(photo []ds.Photo) error
	SaveAudio(audio ds.Audio) error
	GetMessages(userId, chatId uint) ([]MessageTemp, error)
	ChangeChatUserAdmin(chatId uint) error
	GetChat(chatId uint) (ChatInp, error)
	GetAttachments(messId uint) ([]string, error)
	GetPhotos(messId uint) ([]string, error)
	CheckReaction(reaction ds.Reaction) (bool, error)
	CreateReaction(reaction ds.Reaction) error
	ChangeReaction(reaction ds.Reaction) error
	GetReactions(messageId uint) ([]ReactionList, error)
	GetReactionGroup(messId uint) ([]ReactionMes, error)
	GetInfoUser(userId uint) (UserFrom, error)
}
