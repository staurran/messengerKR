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

func (uc *UserUseCase) DeleteContact(userId uint, contactId uint) error {
	err := uc.UserRepo.DeleteContact(userId, contactId)
	return err
}

func (uc *UserUseCase) GetContacts(userId uint) ([]user.Contact, error) {
	contacts, err := uc.UserRepo.GetAllContacts(userId)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func (uc *UserUseCase) GetUserById(userId uint) (user.UserInfo, error) {
	userInfo, err := uc.UserRepo.GetUserById(userId)
	if err != nil {
		return user.UserInfo{}, err
	}
	return userInfo, nil
}
