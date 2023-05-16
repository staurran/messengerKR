package chat

import "time"

type ChatInp struct {
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Users       []uint `json:"users"`
}

type ChatRepoStruct struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	CountMes int64  `json:"count_mes"`
}

type ChatStruct struct {
	Id                  uint   `json:"id"`
	Name                string `json:"name"`
	Avatar              string `json:"avatar"`
	QuantityUnshownMess int64  `json:"quantityUnshownMess"`
	LastMessage
}

type LastMessage struct {
	Username string    `json:"username"`
	Content  string    `json:"content"`
	Time     time.Time `json:"time"`
}

type Message struct {
	ID           uint          `json:"id"`
	Content      string        `json:"content"`
	UserFrom     string        `json:"userFrom"`
	OriginalUser string        `json:"originalUser"`
	Attachment   []string      `json:"attachment"`
	Photos       []string      `json:"photo"`
	Audio        string        `json:"audio"`
	TimeCreated  time.Time     `json:"timeCreated"`
	Reactions    []ReactionMes `json:"reactions"`
	Shown        bool          `json:"shown"`
}

type ReactionMes struct {
	Reaction string `json:"reaction"`
	Quantity int    `json:"quantity"`
}

type MessageInp struct {
	Content    string   `json:"content"`
	ChatId     uint     `json:"chatId"`
	Attachment []string `json:"attachment"`
	Photos     []string `json:"photos"`
	Audio      string   `json:"audio"`
}
