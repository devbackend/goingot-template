package helper

import "github.com/devbackend/goingot/pkg/logger"

type FnError func() error

func DeferError(comment string, logger logger.Logger, fn FnError) {
	err := fn()
	if err != nil {
		logger.Error(comment, err)
	}
}
