package service

import (
	"errors"
	"strings"
)

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type StringSvc struct{}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")

func (StringSvc) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (StringSvc) Count(s string) int {
	return len(s)
}
