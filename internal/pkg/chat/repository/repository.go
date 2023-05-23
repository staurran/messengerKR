package repository

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/staurran/messengerKR.git/internal/app/constProject"
	"github.com/staurran/messengerKR.git/internal/app/ds"
	"github.com/staurran/messengerKR.git/internal/pkg/chat"
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

func (r *ChatRepository) SaveChatUsers(user []ds.ChatUser) error {
	err := r.db.Create(&user).Error
	return err
}

func (r *ChatRepository) DeleteChatUser(user *ds.ChatUser) error {

	err := r.db.First(user, "user_id =? AND chat_id =?", user.UserID, user.ChatID).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&ds.ChatUser{}, "user_id =? AND chat_id =?", user.UserID, user.ChatID).Error
	return err
}

func (r *ChatRepository) GetChats(userID uint) ([]chat.ChatRepoStruct, error) {
	var chatId []uint
	err := r.db.Table("chat_users").Select("chat_id").Find(&chatId, "user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	var chat []chat.ChatRepoStruct
	err = r.db.Table("chats ch").Select("ch.id, ch.name, ch.avatar, COUNT(m.id) as count_mes").
		Joins("Left Join messages m ON m.chat_id = ch.id").
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

func (r *ChatRepository) GetLastMes(chatId uint) (lastMessage chat.LastMessage, err error) {
	err = r.db.Table("messages m").Select("m.context as content, u.username").
		Joins("Join users u ON u.id = m.user_from_id").
		Where("m.chat_id = ?", chatId).
		Order("time_created desc").
		Limit(1).
		Find(&lastMessage).Error

	return lastMessage, nil
}

func (r *ChatRepository) CreateMessage(message *ds.Message) error {
	message.Time = time.Now()
	err := r.db.Create(message).Error
	return err
}

func (r *ChatRepository) CreateMesUserShown(msg []ds.Shown) error {
	err := r.db.Create(&msg).Error
	return err
}

func (r *ChatRepository) GetChatUsers(chatId uint) ([]uint, error) {
	var chatusers []uint
	err := r.db.Table("chat_users").Select("user_id").
		Find(&chatusers, "chat_id = ?", chatId).Error
	return chatusers, err
}

func (r *ChatRepository) SaveAttachments(attachment []ds.Attachment) error {
	err := r.db.Create(&attachment).Error
	return err
}

func (r *ChatRepository) SavePhoto(photos []ds.Photo) error {
	err := r.db.Create(&photos).Error
	return err
}

func (r *ChatRepository) SaveAudio(audio ds.Audio) error {
	err := r.db.Create(&audio).Error
	return err
}

func (r *ChatRepository) GetMessages(userId, chatId uint) ([]chat.MessageTemp, error) {
	var messages []chat.MessageTemp
	err := r.db.Table("messages m").
		Select("m.id, m.content, m.user_id, m.time as time_created").
		Joins("Left Join audios a on a.message_id=m.id").
		//Joins("Join showns s on s.message_id=m.id").
		//Where("s.user_id = ?", userId).
		Where("m.chat_id = ?", chatId).
		Find(&messages).Error
	return messages, err
}

func (r *ChatRepository) ChangeChatUserAdmin(chatId uint) error {
	var randUser ds.ChatUser
	err := r.db.First(&randUser, "chat_id =?", chatId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}

	randUser.ChatRole = constProject.ChatAdmin
	err = r.db.Save(&randUser).Error
	return err
}

func (r *ChatRepository) GetChat(chatId uint) (chat.ChatInp, error) {
	var chatDB chat.ChatInp
	err := r.db.Table("chats c").Select("c. id, c.name, c.description, c.avatar").
		Where("c.id=?", chatId).Find(&chatDB).Error
	return chatDB, err
}

func (r *ChatRepository) GetAttachments(messId uint) ([]string, error) {
	var attachments []string
	err := r.db.Table("attachments").Select("attachment").Where("message_id = ?", messId).Find(&attachments).Error
	return attachments, err
}

func (r *ChatRepository) GetPhotos(messId uint) ([]string, error) {
	var photos []string
	err := r.db.Table("photos").Select("photo").Where("message_id = ?", messId).Find(&photos).Error
	return photos, err
}

func (r *ChatRepository) CheckReaction(reaction ds.Reaction) (bool, error) {
	var reactionDB ds.Reaction
	err := r.db.First(&reactionDB, "message_id=? AND user_id=?", reaction.MessageID, reaction.UserID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if err == gorm.ErrRecordNotFound {
		return false, nil
	}

	return true, nil
}

func (r *ChatRepository) CreateReaction(reaction ds.Reaction) error {
	err := r.db.Create(&reaction).Error
	return err
}

func (r *ChatRepository) ChangeReaction(reaction ds.Reaction) error {
	var reactionDB ds.Reaction
	err := r.db.Table("reactions").
		Where("message_id=? AND user_id=?", reaction.MessageID, reaction.UserID).
		Find(&reactionDB).Error
	if err != nil {
		return err
	}
	reactionDB.Name = reaction.Name
	err = r.db.Save(&reactionDB).Error
	return err
}

func (r *ChatRepository) GetReactions(messageId uint) ([]chat.ReactionList, error) {
	var reactions []chat.ReactionList
	err := r.db.Table("reactions r").Select("r.user_id, r.name as reaction, u.username, u.avatar").
		Where("r.message_id = ?", messageId).
		Joins("Join users u on u.id=r.user_id").
		Find(&reactions).Error
	if err != nil {
		return nil, err
	}

	return reactions, nil
}

func (r *ChatRepository) GetReactionGroup(messId uint) ([]chat.ReactionMes, error) {
	var reactions []chat.ReactionMes

	err := r.db.Table("reactions r").
		Select("r.name as reaction, COUNT(r.user_id) as quantity").
		Where("r.message_id =?", messId).
		Group("r.name").
		Find(&reactions).Error
	if err != nil {
		return nil, err
	}

	return reactions, nil
}

func (r *ChatRepository) GetInfoUser(userId uint) (chat.UserFrom, error) {
	var user chat.UserFrom
	err := r.db.Table("users").Select("id as user_id, username, avatar").Where("id=?", userId).Find(&user).Error
	return user, err
}
