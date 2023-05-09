package repository

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/staurran/messengerKR.git/internal/app/constProject"
	"github.com/staurran/messengerKR.git/internal/app/ds"
	"github.com/staurran/messengerKR.git/internal/app/structs"
)

type ChatRepository struct {
	db *gorm.DB
}

func (r *ChatRepository) CreateChat(chat ds.Chat) error {
	err := r.db.Create(&chat).Error
	return err
}

func (r *ChatRepository) SaveChatUser(user ds.ChatUser) error {
	err := r.db.Create(&user).Error
	return err
}

func (r *ChatRepository) DeleteChatUser(user ds.ChatUser) error {

	err := r.db.First(&user, "user_id =? AND chat_id =?", user.UserID, user.ChatID).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&ds.ChatUser{}, "user_id =? AND chat_id =?", user.UserID, user.ChatID).Error
	return err
}

func (r *ChatRepository) GetChats(userId uint) ([]structs.ChatStruct, error) {
	var chats []structs.ChatStruct
	err := r.db.Table("chats c").Select(
		"c.id, c.name, c.avatar, c.description").
		Joins("Join chatuser cu ON c.id = cu.chat_id").
		Where("cu.user_id=?", userId).Find(&chats).Error
	return chats, err
}

func (r *ChatRepository) ChangeChat(chatInp ds.Chat) error {
	chatDB := &ds.Chat{}
	err := r.db.First(chatDB, "id = ?", chatInp.Id).Error // find product with code D42
	if err != nil {
		return err
	}
	if chatInp.Name != "" {
		chatDB.Name = chatInp.Name
	}
	if chatInp.Avatar != "" {
		chatDB.Avatar = chatInp.Avatar
	}
	if chatInp.Description != "" {
		chatDB.Description = chatInp.Description
	}
	err = r.db.Save(&chatDB).Error
	return err
}

func (r *ChatRepository) CheckAdmin(userId, chatId uint) error {
	chatUser := &ds.ChatUser{}
	err := r.db.First(chatUser, "user_id = ? AND chat_id = ?", userId, chatId).Error
	if err != nil {
		return err
	}
	if chatUser.ChatRole == constProject.ChatAdmin {
		return nil
	}
	err = fmt.Errorf("user is not an admin")
	return err
}
