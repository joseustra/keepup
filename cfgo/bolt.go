package cfgo

import (
	"encoding/json"
	"log"
	"os"

	"github.com/boltdb/bolt"
)

// Bolt uses boltdb as storage
type Bolt struct {
	DB *bolt.DB
}

// NewBolt returns a new instance of Bolt with the storage initialized
func NewBolt(path string) *Bolt {
	return &Bolt{openBoltDB(path)}
}

// Find finds a domain record with the given key
func (b *Bolt) Find(key string) (*Domain, error) {
	domain := &Domain{}

	err := b.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("domains"))
		v := b.Get([]byte(key))

		if len(v) > 0 {
			json.Unmarshal(v, domain)
		}
		return nil
	})

	return domain, err
}

// Save saves a new domain with the given values
func (b *Bolt) Save(domain *Domain) error {
	err := b.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("domains"))
		j, _ := json.Marshal(domain)
		err := b.Put([]byte(domain.DNS), j)
		return err
	})

	return err
}

func openBoltDB(path string) *bolt.DB {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}

	db, err := bolt.Open(path+"keepup.db", 0755, nil)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	_, err = tx.CreateBucketIfNotExists([]byte("domains"))
	if err != nil {
		log.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}

	return db
}
