package dto

import (
	"time"
)

type Review struct {
	ReviewId    string    `json:"reviewId"`
	UserId      string    `json:"userId"`
	BookTitle   string    `json:"bookTitle"`
	ReviewTitle string    `json:"reviewTitle"`
	Publisher   string    `json:"publisher"`
	Review      string    `json:"review"`
	ReadedAt    time.Time `json:"readedAt"`
	Stars       int       `json:"stars"`
	PublicFlg   bool      `json:"publicFlg"`
	Like        int       `json:"like"`
}

type Reviews struct {
	Reviews []Review `json:"reviews"`
	Page    *Page    `json:"page"`
}

type ReviewWithComments struct {
	ReviewId    string    `json:"reviewId"`
	UserId      string    `json:"userId"`
	BookTitle   string    `json:"bookTitle"`
	ReviewTitle string    `json:"reviewTitle"`
	Publisher   string    `json:"publisher"`
	Review      string    `json:"review"`
	ReadedAt    time.Time `json:"readedAt"`
	Stars       int       `json:"stars"`
	PublicFlg   bool      `json:"publicFlg"`
	Like        int       `json:"like"`
	Comments    Comments  `json:"comments"`
}
