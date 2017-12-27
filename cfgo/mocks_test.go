package cfgo_test

import (
	"github.com/ustrajunior/keepup/cfgo"
)

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
