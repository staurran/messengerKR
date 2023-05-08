package repository

import (
	"github.com/staurran/messengerKR.git/internal/app/structs"
)

func (r *Repository) GetChatMessages(userId, chatId uint) ([]structs.Message, error) {
	var messages []structs.Message
	err := r.db.Table("messages m").Select(
		"m.id, m.context as content, m.time_created, m.user_from_id, m.origin_user_id, "+
			"s.shown, p.link as photo, at.link as attachments, a.link as audio").
		Joins("Join showns s ON m.id = s.message_id").
		Joins("LEFT Join photos p ON p.message_id=m.id").
		Joins("LEFT Join attachments at ON at.message_id=m.id").
		Joins("LEFT Join audios a ON a.message_id=m.id").
		Where("s.user_id=?", userId).
		Where("m.chat_id = ?", chatId).Find(&messages).Error
	return messages, err
}
