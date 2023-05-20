package models

const RootUserName = "root"

type LoginModel struct {
	Account      string `json:"username"`
	OrigPassword string `json:"password"`
}
