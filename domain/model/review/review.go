package review

import (
	"fmt"
	"time"
	"github.com/google/uuid"
	"app/domain/model/user"
)

type Review struct {
	reviewId    ReviewId
	userId      user.UserId
	bookTitle   bookTitle
	reviewTitle reviewTitle
	publisher   publisher
	review      reviewVal
	readedAt    readedAt
	stars       stars
	publicFlg   publicFlg
}

func New(reviewId string, userId string, bookTitle string, reviewTitle string, publisher string, reviewVal string, readedAt time.Time, stars int, publicFlg bool) (*Review, error) {

	createdReviewId, err := NewReviewId(reviewId)
	if err != nil {
		return nil, err
	}

	createdUserId, err := user.NewUserId(userId)
	if err != nil {
		return nil, err
	}

	createdBookTitle, err := newBookTitle(bookTitle)
	if err != nil {
		return nil, err
	}

	createdReviewTitle, err := newReviewTitle(reviewTitle)
	if err != nil {
		return nil, err
	}

	createdPublisher, err := newPublisher(publisher)
	if err != nil {
		return nil, err
	}

	createdReview, err := newReviewComment(reviewVal)
	if err != nil {
		return nil, err
	}

	createdReadedAt, err := newReadedAt(readedAt)
	if err != nil {
		return nil, err
	}

	createdStars, err := newStars(stars)
	if err != nil {
		return nil, err
	}

	createdPublicFlg, err := newPublicFlg(publicFlg)
	if err != nil {
		return nil, err
	}

	review := Review{
		reviewId:    *createdReviewId,
		userId:      *createdUserId,
		bookTitle:   *createdBookTitle,
		reviewTitle: *createdReviewTitle,
		publisher:   *createdPublisher,
		review:      *createdReview,
		readedAt:    *createdReadedAt,
		stars:       *createdStars,
		publicFlg:   *createdPublicFlg,
	}
	return &review, nil
}

func Create(userId string, bookTitle string, reviewTitle string, publisher string, reviewVal string, readedAt time.Time, stars int, publicFlg bool) (*Review, error) {
	reviewId := uuid.New().String()
	review, err := New(reviewId, userId, bookTitle, reviewTitle, publisher, reviewVal, readedAt, stars, publicFlg)

	if err != nil {
		return nil, err
	}

	return review, err
}

func (r Review) GetReviewId() ReviewId {
	return r.reviewId
}

func (r Review) GetUserId() user.UserId {
	return r.userId
}

func (r Review) GetBookTitle() bookTitle {
	return r.bookTitle
}

func (r Review) GetReviewTitle() reviewTitle {
	return r.reviewTitle
}

func (r Review) GetPublisher() publisher {
	return r.publisher
}

func (r Review) GetReview() reviewVal {
	return r.review
}

func (r Review) GetReadedAt() readedAt {
	return r.readedAt
}

func (r Review) GetStars() stars {
	return r.stars
}

func (r Review) GetPublicFlg() publicFlg {
	return r.publicFlg
}

func (r *Review) ChangePublicFlg(current_user_id string) (*Review, error) {

	if current_user_id != string(r.userId) {
		err := fmt.Errorf("%s", "can not change public flg. because this Review is not your review.")
		return nil, err
	}

	if r.publicFlg == true {
		r.publicFlg = false
	} else {
		r.publicFlg = true
	}

	return r, nil
}
