package cfgo_test

import (
	"errors"

	"github.com/ustrajunior/keepup/cfgo"
)

// Mock for success
type MockSuccessStorage struct{}

func (m *MockSuccessStorage) Find(key string) (*cfgo.Domain, error) {
	domain := &cfgo.Domain{
		Zone: "test.com",
		DNS:  "sub.test.com",
		IP:   "127.0.0.1",
	}

	return domain, nil
}

func (m *MockSuccessStorage) Save(domain *cfgo.Domain) error {
	return nil
}

// Mock for error
type MockErrorStorage struct{}

func (m *MockErrorStorage) Find(key string) (*cfgo.Domain, error) {
	return &cfgo.Domain{}, errors.New("Some error on reading")
}

func (m *MockErrorStorage) Save(domain *cfgo.Domain) error {
	return errors.New("Error to save")
}

type MockSuccessClient struct {
	TestDomain *cfgo.Domain
}

func (m *MockSuccessClient) GetDNSRecord(zone string, dns string) (*cfgo.Domain, error) {
	if m.TestDomain == nil {
		return &cfgo.Domain{
			Zone: zone,
			DNS:  dns,
			IP:   "127.0.0.1",
		}, nil
	}

	return m.TestDomain, nil
}

func (m *MockSuccessClient) UpdateDNSRecord(domain *cfgo.Domain) error {
	m.TestDomain = domain
	return nil
}
