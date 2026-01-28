package domain

import "errors"

type ISaltStore interface {
	Create() error
	Exists() (bool, error)
	Get() ([]byte, error)
}

var (
	ErrSaltIsNotExists = errors.New("Salt is not exists. Please run init command first!")
)
