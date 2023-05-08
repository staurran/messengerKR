package photo

import dataStruct "github.com/staurran/messengerKR.git/internal/app/ds"

type IRepositoryPhoto interface {
	SavePhoto(row dataStruct.UserPhoto) error
	DeletePhoto(row dataStruct.UserPhoto) error
	ChangePhoto(photoId, userId, newPhotoId uint) error

	GetAvatar(userId uint) (uint, error)
	GetPhotos(userId uint) ([]uint, error)
}
