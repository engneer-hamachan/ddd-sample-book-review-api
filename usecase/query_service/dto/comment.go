package dto

import (
	"time"
)

type Comment struct {
	CommentId string
	UserId    string
	UserName  string
	Comment   string
	CreatedAt time.Time
}

type Comments struct {
	Comments []Comment `json:"comments"`
	Page     *Page     `json:"page"`
}
