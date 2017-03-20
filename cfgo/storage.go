package cfgo

// Storage defines how a storage should be
type Storage interface {
	Find(key string) (*Domain, error)
	Save(domain *Domain) error
}
