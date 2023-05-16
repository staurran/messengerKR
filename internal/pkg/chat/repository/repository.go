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

func NewChatRepo(db *gorm.DB) *ChatRepository {

	r := ChatRepository{db}
	return &r
}

func (r *ChatRepository) CreateChat(chat *ds.Chat) error {
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

func (r *ChatRepository) GetChats(userID uint, k int) ([]structs.ChatRepoStruct, error) {
	var chatId []uint
	err := r.db.Table("chat_users").Select("chat_id").Find(&chatId, "user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	var chat []structs.ChatRepoStruct
	err = r.db.Table("chats ch").Select("ch.id, ch.name, ch.avatar, COUNT(m.id) as count_mes").
		Joins("Join messages m ON m.chat_id = ch.id").
		Where("ch.id IN ?", chatId).
		Order("COUNT(m.id)").
		Group("ch.id, ch.name, ch.avatar").
		Find(&chat).
		Error
	if err != nil {
		return nil, err
	}
	return chat, nil
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
	if chatInp.Avatar != 0 {
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

func (r *ChatRepository) GetLastMes(chatId uint) (lastMessage structs.LastMessage, err error) {
	err = r.db.Table("messages m").Select("m.context as content, u.username").
		Joins("Join users u ON u.id = m.user_from_id").
		Where("m.chat_id = ?", chatId).
		Order("time_created desc").
		Limit(1).
		Find(&lastMessage).Error

	return lastMessage, nil
}

func (r *ChatRepository) CreateMessage(message *ds.Message) error {
	err := r.db.Create(message).Error
	return err
}

func (r *ChatRepository) CreateMesUserShown(msg ds.Shown) error {
	err := r.db.Create(&msg).Error
	return err
}

func (r *ChatRepository) GetChatUsers(chatId uint) ([]uint, error) {
	var chatusers []uint
	err := r.db.Table("chatusers").Select("user_id").
		Find(&chatusers, "chat_id = ?", chatId).Error
	return chatusers, err
}

func (r *ChatRepository) SaveAttachments(attachment ds.Attachment) error {
	err := r.db.Create(&attachment).Error
	return err
}

func (r *ChatRepository) SavePhotos(photos ds.Photo) error {
	err := r.db.Create(&photos).Error
	return err
}

func (r *ChatRepository) SaveAudio(audio ds.Audio) error {
	err := r.db.Create(&audio).Error
	return err
}
