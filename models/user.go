package models

type Users struct {
	Username		string		`json:"userName"`
	Email			string		`json:"email"`
	Password		string		`json:"password"`
}