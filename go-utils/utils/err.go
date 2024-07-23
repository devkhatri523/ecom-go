package utils

import (
	"errors"
	"fmt"
)

type Error struct {
	ErrorCode string
	Err       error
}

func NewError(errorCode string, err error) Error {
	return Error{
		ErrorCode: errorCode,
		Err:       err,
	}
}

func NewErrorFromStr(errorCode string, err string) Error {
	return Error{
		ErrorCode: errorCode,
		Err:       errors.New(err),
	}
}

func NewErrorFromPanic(errorCode string, err any) Error {
	return NewError(errorCode, errors.New(fmt.Sprintf("%s", err)))
}

func (m Error) Error() string {
	return fmt.Sprintf("%s : %s", m.ErrorCode, m.Err)
}
