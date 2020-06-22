package helper

import (
	"errors"
	"github.com/devbackend/goingot/pkg/logger/logger_mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestDeferError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	comment := "test comment"
	err := errors.New("test error")

	mockLogger := logger_mock.NewMockLogger(ctl)
	mockLogger.EXPECT().Error(comment, err)

	DeferError("test comment", mockLogger, func() error {
		return err
	})
}
