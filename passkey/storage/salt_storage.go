package storage

import (
	"crypto/rand"
	"os"
	"sync"

	"github.com/Efesngl/learngo/passkey/domain"
)

type SaltStorage struct {
	path string
	mu   sync.Mutex
}

func NewSaltStorage(path string) *SaltStorage {
	return &SaltStorage{
		path: path,
	}
}
func (s *SaltStorage) Exists() (bool, error) {
	_, err := os.Stat(s.path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (s *SaltStorage) Create() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	salt := make([]byte, 32)
	rand.Read(salt)
	err := os.WriteFile(s.path, salt, 0600)
	if err != nil {
		return err
	}
	return nil
}
func (s *SaltStorage) Get() ([]byte, error) {
	exists, err := s.Exists()
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, domain.ErrSaltIsNotExists
	}
	
	file, err := os.ReadFile(s.path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
