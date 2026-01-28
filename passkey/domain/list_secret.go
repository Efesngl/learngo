package domain

type ListSecrets struct {
	store ISecretStore
}

func NewListSecrets(store ISecretStore) *ListSecrets {
	return &ListSecrets{
		store: store,
	}
}
func (ls *ListSecrets) Execute() ([]Secret, error) {
	return ls.store.List()
}
