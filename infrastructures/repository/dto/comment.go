package dto

import (
	"app/domain/model/review/comment"
	"time"
)

type Comment struct {
	ID        int
	CommentId string
	ReviewId  string
	UserId    string
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ConvertComment(c *comment.Comment) *Comment {
	return &Comment{
		CommentId: string(c.GetCommentId()),
		ReviewId:  string(c.GetReviewId()),
		UserId:    string(c.GetUserId()),
		Comment:   string(c.GetComment()),
	}
}

func AdaptComment(converted_comment *Comment) (*comment.Comment, error) {

	comment, err := comment.New(converted_comment.CommentId, converted_comment.ReviewId, converted_comment.UserId, converted_comment.Comment)

	if err != nil {
		return nil, err
	}
	return comment, nil
}
