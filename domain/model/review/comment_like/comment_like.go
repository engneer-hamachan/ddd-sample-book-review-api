package comment_like

import (
	"app/domain/model/vo"
	"fmt"
	"github.com/google/uuid"
)

type CommentLike struct {
	commentLikeId vo.UuId
	commentId     vo.UuId
	userId        vo.UuId
}

func New(commentLikeId string, commentId string, userId string) (*CommentLike, error) {

	createdCommentLikeId, err := vo.NewUuId(commentLikeId)
	if err != nil {
		return nil, err
	}

	createdCommentId, err := vo.NewUuId(commentId)
	if err != nil {
		return nil, err
	}

	createdUserId, err := vo.NewUuId(userId)
	if err != nil {
		return nil, err
	}

	commentLike := CommentLike{
		commentLikeId: *createdCommentLikeId,
		commentId:     *createdCommentId,
		userId:        *createdUserId,
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

func (c CommentLike) GetCommentLikeLikeId() vo.UuId {
	return c.commentLikeId
}

func (c CommentLike) GetCommentId() vo.UuId {
	return c.commentId
}

func (c CommentLike) GetUserId() vo.UuId {
	return c.userId
}
