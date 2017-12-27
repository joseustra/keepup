package cfgo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ustrajunior/keepup/cfgo"
)

func TestUpdateInvalidIP(t *testing.T) {
	client := &MockSuccessClient{}

	err := cfgo.UpdateRecord(client, "test.com", "sub.test.com", "127.1.1")
	assert.Error(t, err)
}

func TestUpdateDNS(t *testing.T) {
	client := &MockSuccessClient{}

	err := cfgo.UpdateRecord(client, "test.com", "sub", "127.0.0.1")
	assert.NoError(t, err)

	domain, _ := client.GetDNSRecord("test.com", "sub.test.com")

	assert.Equal(t, "test.com", domain.Zone)
	assert.Equal(t, "sub.test.com", domain.DNS)
	assert.Equal(t, "127.0.0.1", domain.IP)
}

func TestUpdateRecordWithCustomIP(t *testing.T) {
	client := &MockSuccessClient{}

	t.Run("same ip", func(t *testing.T) {
		err := cfgo.UpdateRecord(client, "test.com", "sub.test.com", "127.0.0.1")
		assert.NoError(t, err)

		domain, _ := client.GetDNSRecord("test.com", "sub.test.com")

		assert.Equal(t, "test.com", domain.Zone)
		assert.Equal(t, "sub.test.com", domain.DNS)
		assert.Equal(t, "127.0.0.1", domain.IP)
	})

	t.Run("other ip", func(t *testing.T) {
		err := cfgo.UpdateRecord(client, "test.com", "sub.test.com", "127.0.0.2")
		assert.NoError(t, err)

		domain, _ := client.GetDNSRecord("test.com", "sub.test.com")

		assert.Equal(t, "test.com", domain.Zone)
		assert.Equal(t, "sub.test.com", domain.DNS)
		assert.Equal(t, "127.0.0.2", domain.IP)
	})
}
