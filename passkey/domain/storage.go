package domain

type ISecretStore interface {
	List() ([]Secret, error)
	Save(secret Secret) error
	Exists(name string) (bool, error)
	Get(name string) (Secret, error)
	Delete(name string) error
	DeleteAll() error
}
