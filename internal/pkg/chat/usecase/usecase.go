package usecase

import (
	"time"

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
	admin := ds.ChatUser{UserID: userId, ChatID: chatDB.Id, ChatRole: constProject.ChatAdmin}
	chatUsers = append(chatUsers, admin)
	for _, user := range chatInp.Users {
		chatUser := ds.ChatUser{UserID: user, ChatID: chatDB.Id, ChatRole: constProject.ChatUser}
		chatUsers = append(chatUsers, chatUser)
	}
	err = uc.ChatRepo.SaveChatUsers(chatUsers)
	if err != nil {
		return 0, err
	}

	return chatDB.Id, err
}

func (uc *ChatUseCase) DeleteChat(chatId uint, userId uint) error {
	chatUser := ds.ChatUser{UserID: userId, ChatID: chatId}
	err := uc.ChatRepo.DeleteChatUser(chatUser)
	return err
}

func (uc *ChatUseCase) ChangeChat(chatId uint, userId uint, chat chat.ChatInp) error {
	err := uc.ChatRepo.CheckAdmin(userId, chatId)
	if err != nil {
		return err
	}

	NewChat := ds.Chat{Name: chat.Name, Description: chat.Description}
	err = uc.ChatRepo.ChangeChat(NewChat)
	return err
}

func (uc *ChatUseCase) GetChats(userId uint) ([]chat.ChatStruct, error) {
	chats, err := uc.ChatRepo.GetChats(userId)
	if err != nil {
		return nil, err
	}

	var result []chat.ChatStruct
	for _, c := range chats {
		lastMes, err := uc.ChatRepo.GetLastMes(c.Id)
		if err != nil {
			return nil, err
		}
		resultChat := chat.ChatStruct{c.Id, c.Name, c.Avatar, c.CountMes, lastMes}

		result = append(result, resultChat)
	}
	return result, nil
}

func (uc *ChatUseCase) CreateMessage(userId uint, inp chat.MessageInp) error {
	msg := ds.Message{Content: inp.Content, UserId: userId, ChatId: inp.ChatId, Time: time.Now(), Shown: false}
	chatUsers, err := uc.ChatRepo.GetChatUsers(inp.ChatId)
	if err != nil {
		return err
	}
	err = uc.ChatRepo.CreateMessage(&msg)
	if err != nil {
		return err
	}

	var shownSlice []ds.Shown
	for _, user := range chatUsers {
		showRow := ds.Shown{UserId: user, MessageId: msg.Id, Shown: false}
		shownSlice = append(shownSlice, showRow)
	}
	err = uc.ChatRepo.CreateMesUserShown(shownSlice)
	if err != nil {
		return err
	}

	var attachmentSlice []ds.Attachment
	for _, attachment := range inp.Attachment {
		attachRow := ds.Attachment{MessageID: msg.Id, Attachment: attachment}
		attachmentSlice = append(attachmentSlice, attachRow)
	}
	err = uc.ChatRepo.SaveAttachments(attachmentSlice)
	if err != nil {
		return err
	}

	var photoSlice []ds.Photo
	for _, photo := range inp.Photos {
		photoRow := ds.Photo{MessageID: msg.Id, Photo: photo}
		attachmentSlice = append(attachmentSlice, photoRow)
	}
	err = uc.ChatRepo.SaveAttachments(attachmentSlice)
	if err != nil {
		return err
	}

	if inp.Audio != 0 {
		audio := ds.Audio{
			Audio:     inp.Audio,
			MessageID: msg.Id,
		}
	}

	return nil
}
