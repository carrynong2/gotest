package services

import "errors"

var (
	ErrZeroAmount = errors.New("purchase amount zero")
	ErrRepository = errors.New("repository error")
)