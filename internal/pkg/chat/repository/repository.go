package repository

import (
	"gorm.io/gorm"

	dataStruct "github.com/staurran/messengerKR.git/internal/app/ds"
	"github.com/staurran/messengerKR.git/internal/app/structs"
)

type ChatRepository struct {
	db *gorm.DB
}

func (r *ChatRepository) CreateChat(chat dataStruct.Chat) error {
	err := r.db.Create(&chat).Error
	return err
}

func (r *ChatRepository) SaveChatUser(user dataStruct.ChatUser) error {
	err := r.db.Create(&user).Error
	return err
}

func (r *ChatRepository) DeleteChatUser(user dataStruct.ChatUser) error {

	err := r.db.First(&user, "user_id =? AND chat_id =?", user.UserId, user.ChatID).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&dataStruct.ChatUser{}, "user_id =? AND chat_id =?", user.UserId, user.Photo).Error
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
