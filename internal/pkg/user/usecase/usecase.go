package usecase

import (
	"github.com/staurran/messengerKR.git/internal/app/ds"
	"github.com/staurran/messengerKR.git/internal/pkg/user"
)

type UserUseCase struct {
	UserRepo user.IUserRepository
}

func (uc *UserUseCase) CreateContact(userId uint, phone string) error {
	contactId, err := uc.UserRepo.GetUserIdByPhone(phone)
	if err != nil {
		return err
	}
	contectRow := ds.Contact{UserID: userId, ContactID: contactId}
	err = uc.UserRepo.CreateContact(contectRow)
	if err != nil {
		return err
	}
	return nil
}
