package domain

import "errors"

type Secret struct {
	Name  string
	Value string
}

var (
	ErrEmptyName  = errors.New("name cannot be empty")
	ErrEmptyValue = errors.New("value cannot be empty")
	ErrDuplicate  = errors.New("secret already exists")
)
