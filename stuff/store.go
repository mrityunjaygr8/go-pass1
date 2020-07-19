package stuff

import (
	"log"

	"github.com/boltdb/bolt"
)

// Store is a struct to hold the BoltDB
type Store struct {
	db *bolt.DB
}

// Init opens a bolt database and returns it so that it can be used
func Init() (Store, error) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return Store{db}, nil
}

// Close closes the database opened in Init
func (s *Store) Close() error {
	err := s.db.Close()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// // Transaction provides a transaction for doing stuff
// func Transaction(s *Store) (bolt.Tx, error) {

// }
