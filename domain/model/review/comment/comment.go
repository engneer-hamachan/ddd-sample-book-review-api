package comment

import (
	"github.com/google/uuid"
	"app/domain/model/review"
	"app/domain/model/user"
)

type Comment struct {
	commentId CommentId
	reviewId  review.ReviewId
	userId    user.UserId
	comment   commentVal
}

func New(commentId string, reviewId string, userId string, commentVal string) (*Comment, error) {

	createdCommentId, err := NewCommentId(commentId)
	if err != nil {
		return nil, err
	}

	createdReviewId, err := review.NewReviewId(reviewId)
	if err != nil {
		return nil, err
	}

	createdUserId, err := user.NewUserId(userId)
	if err != nil {
		return nil, err
	}

	createdComment, err := newCommentVal(commentVal)
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

func (c Comment) GetCommentId() CommentId {
	return c.commentId
}

func (c Comment) GetReviewId() review.ReviewId {
	return c.reviewId
}

func (c Comment) GetUserId() user.UserId {
	return c.userId
}

func (c Comment) GetComment() commentVal {
	return c.comment
}
