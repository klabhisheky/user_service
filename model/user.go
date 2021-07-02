package model

type User struct {
	UserId  int    `json:"UserId,omitempty"`
	Name    string `json:"Name,omitempty"`
	Address string `json:"Address,omitempty"`
	Phone   string `json:"Phone,omitempty"`
}
