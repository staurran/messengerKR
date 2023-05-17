package chat

import "github.com/staurran/messengerKR.git/internal/app/ds"

type UseCase interface {
	CreateChat(chat ChatInp, userId uint) (chatId uint, err error)
	DeleteChat(chatId uint, userId uint) error
	ChangeChat(chatId uint, userId uint, chat ChatInp) error
	GetChats(userId uint) ([]ChatStruct, error)

	GetMessages(userId uint, chatId uint) ([]Message, error)
	CreateMessage(userId uint, inp MessageInp) error

	CreateReaction(reaction ds.Reaction) error
	GetReaction(messageId uint) ([]ReactionList, error)
}
