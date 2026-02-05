package domain

type DeleteAllSecret struct {
	store ISecretStore
}

func NewDeleteAllSecret(store ISecretStore) *DeleteAllSecret {
	return &DeleteAllSecret{store: store}
}

func (as *DeleteAllSecret) Execute() error {
	return as.store.DeleteAll()
}
