package utils

import (
	"context"
	"errors"
)

func HandleContextError(err error) error {
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}
