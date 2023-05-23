package usecase

import (
	"time"

	"github.com/staurran/messengerKR.git/internal/app/constProject"
	"github.com/staurran/messengerKR.git/internal/app/ds"
	"github.com/staurran/messengerKR.git/internal/pkg/chat"
)

type ChatUseCase struct {
	ChatRepo chat.IRepositoryChat
}

func NewChatUseCase(
	chatRepo chat.IRepositoryChat,
) *ChatUseCase {
	return &ChatUseCase{
		ChatRepo: chatRepo,
	}
}

func (uc *ChatUseCase) CreateChat(chatInp chat.ChatInp, userId uint) (chatId uint, err error) {
	chatDB := ds.Chat{Name: chatInp.Name, Description: chatInp.Description, Avatar: chatInp.Avatar}
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
	err := uc.ChatRepo.CheckAdmin(userId, chatId)
	if err != nil && err.Error() != "user is not an admin" {
		return err
	}
	flagNewAdmin := false
	if err == nil {
		flagNewAdmin = true
	}
	err = uc.ChatRepo.DeleteChatUser(&chatUser)
	if flagNewAdmin {
		err2 := uc.ChatRepo.ChangeChatUserAdmin(chatUser.ChatID)
		if err2 != nil {
			return err2
		}
	}
	return err
}

func (uc *ChatUseCase) ChangeChat(chatId uint, userId uint, chat chat.ChatInp) error {
	err := uc.ChatRepo.CheckAdmin(userId, chatId)
	if err != nil {
		return err
	}
	NewChat := ds.Chat{Name: chat.Name, Description: chat.Description, Id: chatId}
	err = uc.ChatRepo.ChangeChat(NewChat)
	if err != nil {
		return err
	}
	chatDB, err := uc.ChatRepo.GetChat(chatId)
	if err != nil {
		return err
	}
	chatUsers, err := uc.ChatRepo.GetChatUsers(chatId)
	if err != nil {
		return err
	}
	chatDB.Users = chatUsers
	var addUsers []ds.ChatUser
	for _, u := range chat.Users {
		if !contains(chatDB.Users, u) {
			addUsers = append(addUsers, ds.ChatUser{UserID: u, ChatID: chatId, ChatRole: constProject.ChatUser})
		}
	}
	if len(addUsers) > 0 {
		err = uc.ChatRepo.SaveChatUsers(addUsers)
		if err != nil {
			return err
		}
	}

	for _, u := range chatDB.Users {
		if !contains(chat.Users, u) {
			err = uc.DeleteChat(chatId, u)
			if err != nil {
				return err
			}
		}
	}

	return nil
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
	msg := ds.Message{Content: inp.Content, UserId: userId, ChatId: inp.ChatId, Time: time.Now()}
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
	if len(attachmentSlice) > 0 {
		err = uc.ChatRepo.SaveAttachments(attachmentSlice)
		if err != nil {
			return err
		}
	}

	var photoSlice []ds.Photo
	for _, photo := range inp.Photos {
		photoRow := ds.Photo{MessageID: msg.Id, Photo: photo}
		photoSlice = append(photoSlice, photoRow)
	}
	if len(photoSlice) > 0 {
		err = uc.ChatRepo.SavePhoto(photoSlice)
		if err != nil {
			return err
		}
	}

	var audio ds.Audio
	if inp.Audio != "" {
		audio = ds.Audio{
			Audio:     inp.Audio,
			MessageID: msg.Id,
		}
		err = uc.ChatRepo.SaveAudio(audio)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc *ChatUseCase) GetMessages(userId uint, chatId uint) ([]chat.Message, error) {
	messages, err := uc.ChatRepo.GetMessages(userId, chatId)
	var result []chat.Message
	if err != nil {
		return nil, err
	}
	for i, m := range messages {

		attachments, err := uc.ChatRepo.GetAttachments(m.ID)
		if err != nil {
			return nil, err
		}
		messages[i].Attachment = attachments

		photos, err := uc.ChatRepo.GetPhotos(m.ID)
		if err != nil {
			return nil, err
		}
		messages[i].Photos = photos

		reactions, err := uc.ChatRepo.GetReactionGroup(m.ID)
		if err != nil {
			return nil, err
		}
		messages[i].Reactions = reactions

		user, err := uc.ChatRepo.GetInfoUser(m.UserId)
		if err != nil {
			return nil, err
		}
		msg := chat.Message{
			ID:          m.ID,
			Content:     m.Content,
			UserFrom:    user,
			Attachment:  attachments,
			Reactions:   reactions,
			Photos:      photos,
			Audio:       m.Audio,
			TimeCreated: m.TimeCreated,
			Shown:       m.Shown,
		}
		result = append(result, msg)
	}

	return result, nil
}

func (uc *ChatUseCase) CreateReaction(reaction ds.Reaction) error {
	exist, err := uc.ChatRepo.CheckReaction(reaction)
	if err != nil {
		return err
	}
	if !exist {
		err = uc.ChatRepo.CreateReaction(reaction)
		if err != nil {
			return err
		}
	} else {
		err = uc.ChatRepo.ChangeReaction(reaction)
		if err != nil {
			return err
		}
	}
	return nil
}

func (uc *ChatUseCase) GetReaction(messageId uint) ([]chat.ReactionList, error) {
	reactions, err := uc.ChatRepo.GetReactions(messageId)
	if err != nil {
		return nil, err
	}

	return reactions, err
}

func contains(sliceUint []uint, elem uint) bool {
	for _, i := range sliceUint {
		if i == elem {
			return true
		}
	}
	return false
}
