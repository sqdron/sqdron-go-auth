package model

type User struct  {
	Entity
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}