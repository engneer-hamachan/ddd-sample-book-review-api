package comment_like

import (
	"fmt"
	"github.com/google/uuid"
)

type CommentLike struct {
	commentLikeId commentLikeId
  commentId     commentId
	userId        userId
}

type commentLikeId string
type commentId string
type userId string

func New(commentLikeId string, commentId string, userId string) (*CommentLike, error) {

  createdCommentLikeId, err := newCommentLikeId(commentLikeId)
	if err != nil {
		return nil, err
	}

  createdCommentId, err := newCommentId(commentId)
	if err != nil {
		return nil, err
	}

	createdUserId, err := newUserId(userId)
	if err != nil {
		return nil, err
	}

	commentLike := CommentLike{
		commentLikeId: *createdCommentLikeId,
    commentId:     *createdCommentId,
		userId:       *createdUserId,
	}
	return &commentLike, nil
}

func Create(commentId string, userId string) (*CommentLike, error) {
	commentLikeId := uuid.New().String()
	comment, err := New(commentLikeId, commentId, userId)

	if err != nil {
		return nil, err
	}

	return comment, err
}

func (c CommentLike) GetCommentLikeLikeId() commentLikeId {
	return c.commentLikeId
}

func (c CommentLike) GetCommentId() commentId {
	return c.commentId
}

func (c CommentLike) GetUserId() userId {
	return c.userId
}

func newCommentLikeId(value string) (*commentLikeId, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:commentLikeId newCommentLikeId()")
		return nil, err
	}

	commentLikeId := commentLikeId(value)

	return &commentLikeId, nil
}

func newCommentId(value string) (*commentId, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:commentId newCommentId()")
		return nil, err
	}

	commentId := commentId(value)

	return &commentId, nil
}


func newUserId(value string) (*userId, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:userId newUserId()")
		return nil, err
	}

	userId := userId(value)
	return &userId, nil
}
