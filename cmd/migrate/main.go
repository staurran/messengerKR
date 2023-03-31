package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/staurran/messengerKR.git/internal/app/ds"
	"github.com/staurran/messengerKR.git/internal/app/dsn"
)

func main() {
	_ = godotenv.Load()
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&ds.User{})
	if err != nil {
		panic("cant migrate db goods")
	}

	err = db.AutoMigrate(&ds.Chat{})
	if err != nil {
		panic("cant migrate db goods")
	}

	err = db.AutoMigrate(&ds.Message{})
	if err != nil {
		panic("cant migrate db goods")
	}

	err = db.AutoMigrate(&ds.Reaction{})
	if err != nil {
		panic("cant migrate db goods")
	}

	err = db.AutoMigrate(&ds.ChatUser{})
	if err != nil {
		panic("cant migrate db goods")
	}

	err = db.AutoMigrate(&ds.Contact{})
	if err != nil {
		panic("cant migrate db goods")
	}

	err = db.AutoMigrate(&ds.Audio{})
	if err != nil {
		panic("cant migrate db goods")
	}

	err = db.AutoMigrate(&ds.Photo{})
	if err != nil {
		panic("cant migrate db goods")
	}

	err = db.AutoMigrate(&ds.Attachment{})
	if err != nil {
		panic("cant migrate db goods")
	}

	err = db.AutoMigrate(&ds.Shown{})
	if err != nil {
		panic("cant migrate db goods")
	}
}
