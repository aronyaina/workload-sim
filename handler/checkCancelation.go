package handler

import "errors"

func CheckCancelation(cancel <-chan struct{}) error {
	select {
	case <-cancel:
		return errors.New("request cancelled")
	default:
	}
	return nil
}
