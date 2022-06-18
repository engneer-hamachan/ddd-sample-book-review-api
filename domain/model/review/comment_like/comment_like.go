package comment_like

import (
	"fmt"
	"github.com/google/uuid"
	"app/domain/model/review/comment"
	"app/domain/model/user"
)

type CommentLike struct {
	commentLikeId commentLikeId
  commentId     comment.CommentId
	userId        user.UserId
}

func New(commentLikeId string, commentId string, userId string) (*CommentLike, error) {

  createdCommentLikeId, err := newCommentLikeId(commentLikeId)
	if err != nil {
		return nil, err
	}

  createdCommentId, err := comment.NewCommentId(commentId)
	if err != nil {
		return nil, err
	}

	createdUserId, err := user.NewUserId(userId)
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

func (c CommentLike) GetCommentLikeId() commentLikeId {
	return c.commentLikeId
}

func (c CommentLike) GetCommentId() comment.CommentId {
	return c.commentId
}

func (c CommentLike) GetUserId() user.UserId {
	return c.userId
}

func (c CommentLike) IsCommentLikeYours(current_user_id string) (b bool, err error) {

	if string(c.userId) != current_user_id {
		err := fmt.Errorf("%s", "this comment like is not yours")
		return false, err
	}

	return true, nil
}
