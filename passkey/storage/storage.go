package storage

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/Efesngl/learngo/passkey/domain"
)

type JSONStorage struct {
	path string
	mu   sync.Mutex
}

func NewJSONStorage(path string) *JSONStorage {
	return &JSONStorage{
		path: path,
	}
}

func (s *JSONStorage) load() ([]domain.Secret, error) {
	file, err := os.ReadFile(s.path)
	if os.IsNotExist(err) {
		return []domain.Secret{}, nil
	}
	if err != nil {
		return nil, err
	}
	var secrets []domain.Secret
	if err := json.Unmarshal(file, &secrets); err != nil {
		return nil, err
	}

	return secrets, nil
}
func (s *JSONStorage) Save(secret domain.Secret) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	secrets, err := s.load()
	if err != nil {
		return err
	}
	secrets = append(secrets, secret)
	return s.SaveAll(secrets)
}

func (s *JSONStorage) SaveAll(secrets []domain.Secret) error {
	data, err := json.MarshalIndent(secrets, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0644)
}

func (s *JSONStorage) Exists(name string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	secrets, err := s.load()
	if err != nil {
		return false, err
	}
	for _, secret := range secrets {
		if secret.Name == name {
			return true, nil
		}
	}
	return false, nil
}
func (s *JSONStorage) List() ([]domain.Secret, error) {
	secrets, err := s.load()
	if err != nil {
		return []domain.Secret{}, err
	}
	return secrets, nil
}
func (s *JSONStorage) Get(name string) (domain.Secret, error) {
	secrets, err := s.load()
	if err != nil {
		return domain.Secret{}, err
	}
	for _, secret := range secrets {
		if secret.Name == name {
			return secret, nil
		}
	}
	return domain.Secret{}, domain.ErrSecretIsNotExists
}
func (s *JSONStorage) Delete(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	secrets, err := s.load()
	if err != nil {
		return err
	}
	filtered := make([]domain.Secret, 0)

	for _, secret := range secrets {
		if secret.Name != name {
			filtered = append(filtered, secret)
		}
	}
	return s.SaveAll(filtered)
}
func (s *JSONStorage) DeleteAll() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.SaveAll([]domain.Secret{})
}
func (s *JSONStorage) First() (string, error) {
	secrets, err := s.load()
	if err != nil {
		return "", err
	}
	if len(secrets) == 0 {
		return "", nil
	}
	return secrets[0].Value, nil
}
