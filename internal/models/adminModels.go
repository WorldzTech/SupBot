package models

type Admin struct {
	Id       uint   `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type AdminModelLogpassFilter struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
