package user

import "github.com/staurran/messengerKR.git/internal/app/ds"

type IUserRepository interface {
	GetUserIdByPhone(phone string) (uint, error)
	CreateContact(contact ds.Contact) error
}
