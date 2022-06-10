package dto

import (
	"app/domain/model/review/comment_like"
)

type CommentLike struct {
	CommentLikeId string
  CommentId     string
	UserId       string
}

func ConvertCommentLike(c *comment_like.CommentLike) *CommentLike {
	return &CommentLike{
		CommentLikeId: string(c.GetCommentLikeId()),
		CommentId:  string(c.GetCommentId()),
		UserId:    string(c.GetUserId()),
	}
}

func AdaptCommentLike(converted_comment_like *CommentLike) (*comment_like.CommentLike, error) {

	comment_like, err := comment_like.New(converted_comment_like.CommentLikeId, converted_comment_like.CommentId, converted_comment_like.UserId)

	if err != nil {
		return nil, err
	}
	return comment_like, nil
}
