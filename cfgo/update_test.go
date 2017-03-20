package cfgo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ustrajunior/keepup/cfgo"
)

func TestUpdateRecordSameIP(t *testing.T) {
	storage := &MockSuccessStorage{}
	client := &MockSuccessClient{}

	err := cfgo.UpdateRecord(storage, client, "test.com", "sub.test.com", "127.0.0.1", false)
	assert.Equal(t, cfgo.ErrSameIP, err)
}

func TestUpdateRecordWithCustomIP(t *testing.T) {
	storage := &MockSuccessStorage{}
	client := &MockSuccessClient{}

	err := cfgo.UpdateRecord(storage, client, "test.com", "sub.test.com", "127.0.0.1", true)
	assert.NoError(t, err)

	domain, _ := client.GetDNSRecord("test.com", "sub.test.com")

	assert.Equal(t, "test.com", domain.Zone)
	assert.Equal(t, "sub.test.com", domain.DNS)
	assert.Equal(t, "127.0.0.1", domain.IP)
}

func TestUpdateRecordWithCurrentIPV4(t *testing.T) {
	storage := &MockSuccessStorage{}
	client := &MockSuccessClient{}

	err := cfgo.UpdateRecord(storage, client, "test.com", "sub.test.com", "", true)
	assert.NoError(t, err)

	domain, _ := client.GetDNSRecord("test.com", "sub.test.com")

	ipv4, _ := cfgo.GetIPV4IP()

	assert.Equal(t, "test.com", domain.Zone)
	assert.Equal(t, "sub.test.com", domain.DNS)
	assert.Equal(t, ipv4, domain.IP)
}
