package review

import (
	"app/domain/model/vo"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Review struct {
	reviewId    vo.UuId
	userId      vo.UuId
	bookTitle   vo.Title
	reviewTitle vo.Title
	publisher   vo.Publisher
	review      vo.Review
	readedAt    vo.ReadedAt
	stars       vo.Stars
	publicFlg   vo.Flag
}

type publicFlg bool

func New(reviewId string, userId string, bookTitle string, reviewTitle string, publisher string, reviewVal string, readedAt time.Time, stars int, publicFlg bool) (*Review, error) {

	createdReviewId, err := vo.NewUuId(reviewId)
	if err != nil {
		return nil, err
	}

	createdUserId, err := vo.NewUuId(userId)
	if err != nil {
		return nil, err
	}

	createdBookTitle, err := vo.NewTitle(bookTitle)
	if err != nil {
		return nil, err
	}

	createdReviewTitle, err := vo.NewTitle(reviewTitle)
	if err != nil {
		return nil, err
	}

	createdPublisher, err := vo.NewPublisher(publisher)
	if err != nil {
		return nil, err
	}

	createdReview, err := vo.NewReview(reviewVal)
	if err != nil {
		return nil, err
	}

	createdReadedAt, err := vo.NewReadedAt(readedAt)
	if err != nil {
		return nil, err
	}

	createdStars, err := vo.NewStars(stars)
	if err != nil {
		return nil, err
	}

	createdPublicFlg, err := vo.NewFlag(publicFlg)
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

func Create(userId string,
	bookTitle string,
	reviewTitle string,
	publisher string,
	reviewVal string,
	readedAt time.Time,
	stars int,
	publicFlg bool) (*Review, error) {
	reviewId := uuid.New().String()
	review, err :=
		New(reviewId, userId, bookTitle, reviewTitle, publisher, reviewVal, readedAt, stars, publicFlg)
	if err != nil {
		return nil, err
	}

	return review, err
}

func (r Review) GetReviewId() vo.UuId {
	return r.reviewId
}

func (r Review) GetUserId() vo.UuId {
	return r.userId
}

func (r Review) GetBookTitle() vo.Title {
	return r.bookTitle
}

func (r Review) GetReviewTitle() vo.Title {
	return r.reviewTitle
}

func (r Review) GetPublisher() vo.Publisher {
	return r.publisher
}

func (r Review) GetReview() vo.Review {
	return r.review
}

func (r Review) GetReadedAt() vo.ReadedAt {
	return r.readedAt
}

func (r Review) GetStars() vo.Stars {
	return r.stars
}

func (r Review) GetPublicFlg() vo.Flag {
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
