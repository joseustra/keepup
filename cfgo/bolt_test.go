package cfgo_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/stretchr/testify/assert"
	"github.com/ustrajunior/keepup/cfgo"
)

func newBoltTest() *cfgo.Bolt {
	path := "/tmp/"
	return cfgo.NewBolt(path)
}

func TestSave(t *testing.T) {
	b := newBoltTest()
	defer os.Remove(b.DB.Path())

	d := &cfgo.Domain{
		ZoneID:   "123123",
		RecordID: "321321",
		Type:     "A",
		Zone:     "test.com",
		DNS:      "test.com",
		IP:       "127.0.0.1",
	}

	err := b.Save(d)
	assert.NoError(t, err)

	domain := &cfgo.Domain{}

	b.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("domains"))
		v := b.Get([]byte("test.com"))

		if len(v) > 0 {
			json.Unmarshal(v, domain)
		}
		return nil
	})

	assert.EqualValues(t, d, domain)
}

func TestFind(t *testing.T) {
	b := newBoltTest()
	defer os.Remove(b.DB.Path())

	d := &cfgo.Domain{
		ZoneID:   "123123",
		RecordID: "321321",
		Type:     "A",
		Zone:     "test.com",
		DNS:      "test.com",
		IP:       "127.0.0.1",
	}

	b.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("domains"))
		j, _ := json.Marshal(d)
		err := b.Put([]byte(d.DNS), j)
		return err
	})

	domain, err := b.Find(d.DNS)
	assert.NoError(t, err)
	assert.EqualValues(t, d, domain)
}
