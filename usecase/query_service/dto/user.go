package dto

type User struct {
	UserId     string `json:"userId"`
	Name       string `json:"name"`
	Permission string `json:"permission"`
}

type Users struct {
	Users []User `json:"users"`
	Page  *Page  `json:"page"`
}

type UserForAuth struct {
	UserId   string `json:"userId"`
	Name     string `json:"userName"`
	Mail     string `json:"mailadress"`
	Password string `json:"password"`
}
