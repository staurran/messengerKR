package chat

import "time"

type ChatInp struct {
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Users       []uint `json:"users"`
}

type ChatRepoStruct struct {
	Id       uint `json:"id"`
	Name     string
	Avatar   string
	CountMes int64 `json:"count_mes"`
}

type ChatStruct struct {
	Id                  uint `json:"id"`
	Name                string
	Avatar              string
	QuantityUnshownMess int64 `json:"quantityUnshownMess"`
	LastMessage
}

type LastMessage struct {
	Username string
	Content  string
	Time     time.Time
}

type Message struct {
	ID           uint   `json:"id"`
	Content      string `json:"content"`
	UserFrom     string `json:"userFrom"`
	OriginalUser string `json:"originalUser"`
	Attachment   []uint
	Photos       []uint
	Audio        uint
	TimeCreated  time.Time
	Reactions    []ReactionMes
	Shown        bool
}

type ReactionMes struct {
	Reaction string
	Quantity int
}

type MessageInp struct {
	Content    string
	Attachment []uint
	Photos     []uint
	Audio      uint
}
