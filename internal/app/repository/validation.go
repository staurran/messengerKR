package repository

import "github.com/staurran/messengerKR.git/internal/app/ds"

func (r *Repository) CheckUsername(username string) (err error) {
	err = r.db.Model(&ds.User{}).Where("login = ?", username).Error
	return
}

func (r *Repository) CheckPhone(phone string) (err error) {
	err = r.db.Model(&ds.User{}).Where("phone = ?", phone).Error
	return
}
