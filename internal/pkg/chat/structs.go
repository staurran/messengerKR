package chat

type ChatInp struct {
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Users       []uint `json:"users"`
}
