package model

type User struct  {
	id     string `json:"id" form:"-"`
	profile string
}