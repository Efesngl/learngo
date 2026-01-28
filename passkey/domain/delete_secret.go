package domain

type DeleteSecret struct {
	store ISecretStore
}

func NewDeleteSecret(store ISecretStore) *DeleteSecret {
	return &DeleteSecret{store: store}
}

func (as *DeleteSecret) Execute(name string) error {
	if name == "" {
		return ErrEmptyName
	}

	exists, err := as.store.Exists(name)
	if err != nil {
		return err
	}

	if !exists {
		return ErrSecretIsNotExists
	}
	return as.store.Delete(name)
}
