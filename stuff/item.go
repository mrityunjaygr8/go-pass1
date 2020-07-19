package stuff

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/mrityunjaygr8/go-pass/stuff/store"
)

// Item is a struct representing an URL-username-password pair
type Item struct {
	URL      string
	username string
	password string
}

func (s *store.Store) AddItem(item Item) error {
	return s.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(item.URL))
		if err != nil {
			log.Fatal(err)
		}

		return b.Put([]byte(item.username), []byte(item.password))
	})
}

func (s *store.Store) GetItem(URL, username string) error {
	return s.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(URL))
		v := b.Get([]byte(username))
		fmt.Printf("%s\n", string(v))
		return nil
	})
}
