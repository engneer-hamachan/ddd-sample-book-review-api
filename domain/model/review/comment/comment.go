package comment

import (
	"app/domain/model/vo"
	"github.com/google/uuid"
	"app/domain/model/review"
	"app/domain/model/user"
)

type Comment struct {
	commentId vo.UuId
	reviewId  vo.UuId
	userId    vo.UuId
	comment   vo.Comment
}

func New(commentId string, reviewId string, userId string, commentVal string) (*Comment, error) {

	createdCommentId, err := vo.NewUuId(commentId)
	if err != nil {
		return nil, err
	}

	createdReviewId, err := vo.NewUuId(reviewId)
	if err != nil {
		return nil, err
	}

	createdUserId, err := vo.NewUuId(userId)
	if err != nil {
		return nil, err
	}

	createdComment, err := vo.NewComment(commentVal)
	if err != nil {
		return nil, err
	}

	comment := Comment{
		commentId: *createdCommentId,
		reviewId:  *createdReviewId,
		userId:    *createdUserId,
		comment:   *createdComment,
	}
	return &comment, nil
}

func Create(reviewId string, userId string, commentVal string) (*Comment, error) {
	commentId := uuid.New().String()
	comment, err := New(commentId, reviewId, userId, commentVal)

	if err != nil {
		return nil, err
	}

	return comment, err
}

func (c *Comment) GetCommentId() vo.UuId {
	return c.commentId
}

func (c *Comment) GetReviewId() vo.UuId {
	return c.reviewId
}

func (c *Comment) GetUserId() vo.UuId {
	return c.userId
}

func (c *Comment) GetComment() vo.Comment {
	return c.comment
}
