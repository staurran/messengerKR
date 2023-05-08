package usecase

import (
	"github.com/staurran/messengerKR.git/internal/app/constProject"
	"github.com/staurran/messengerKR.git/internal/app/ds"
	"github.com/staurran/messengerKR.git/internal/pkg/chat"
)

type ChatUseCase struct {
	ChatRepo chat.IRepositoryPhoto
}

func (uc *ChatUseCase) CreateChat(chatInp chat.ChatInp, userId uint) (chatId uint, err error) {
	chatDB := ds.Chat{Name: chatInp.Name, Description: chatInp.Description, Avatar: 0}
	err = uc.ChatRepo.CreateChat(&chatDB)
	if err != nil {
		return 0, err
	}

	var chatUsers []ds.ChatUser
	admin := ds.ChatUser{UserID: userId, ChatID: chatDB.ID, ChatRole: constProject.ChatAdmin}
	chatUsers = append(chatUsers, admin)
	for _, user := range chatInp.Users {
		chatUser := ds.ChatUser{UserID: user, ChatID: chatDB.ID, ChatRole: constProject.ChatUser}
		chatUsers = append(chatUsers, chatUser)
	}
	err = uc.ChatRepo.SaveChatUsers(chatUsers)
	if err != nil {
		return 0, err
	}

	return chatDB.ID, err
}

func (uc *ChatUseCase) DeleteChat(chatId uint, userId uint) error {
	chatUser := ds.ChatUser{UserID: userId, ChatID: chatId}
	err := uc.ChatRepo.DeleteChatUser(chatUser)
	return err
}
