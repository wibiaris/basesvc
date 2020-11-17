package controller

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/iwanjunaid/basesvc/domain/model"

	"github.com/golang/mock/gomock"

	// adapter "github.com/iwanjunaid/basesvc/shared/mock/adapter"

	"github.com/iwanjunaid/basesvc/shared/mock/interactor"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInsertAuthorController(t *testing.T) {
	Convey("Insert Author Controller", t, func() {
		var (
			entAuthor = &model.Author{
				Name:      "123",
				Email:     "123",
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			}
		)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		authorInteractor := interactor.NewMockAuthorInteractor(ctrl)
		Convey("Negative Scenarios", func() {
			Convey("Should return error", func() {
				authorInteractor.EXPECT().Create(context.Background(), entAuthor).Return(nil, errors.New("error"))
				ctrl := NewAuthorController(authorInteractor)
				_, err := ctrl.InsertAuthor(context.Background(), entAuthor)
				So(err, ShouldBeError)
			})
		})
		Convey("Positive Scenarios", func() {
			Convey("Should return error", func() {
				author := &model.Author{
					Name:      "123",
					Email:     "123",
					CreatedAt: time.Now().Unix(),
					UpdatedAt: time.Now().Unix(),
				}
				authorInteractor.EXPECT().Create(context.Background(), author).Return(author, nil)
				ctrl := NewAuthorController(authorInteractor)
				res, _ := ctrl.InsertAuthor(context.Background(), author)
				So(res, ShouldEqual, author)
			})
		})
	})
}
