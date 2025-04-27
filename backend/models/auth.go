package models

type Login struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginDTO struct {
	UserName string `json:"username"`
	Token    string `json:"token"`
}
