package stuff

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
	// "github.com/mrityunjaygr8/go-pass/stuff"
)

// Item is a struct representing an URL-username-password pair
type Item struct {
	URL      string
	Username string
	Password string
}

// AddItem adds a new record to the database
func (s *Store) AddItem(item Item) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(item.URL))
		if err != nil {
			log.Fatal(err)
		}

		return b.Put([]byte(item.Username), []byte(item.Password))
	})
}

// GetItem fetches a given record from the database
func (s *Store) GetItem(URL, username string) error {
	return s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(URL))
		if b == nil {
			fmt.Println("There are no saved credentials for this URL")
			fmt.Println("You can create new credentials using `go-pass add --URL url_name --username user_name --password pass_word`")
			os.Exit(1)
		}
		v := b.Get([]byte(username))
		if v == nil {
			fmt.Println("There are no saved credentials for this username on this URL")
			fmt.Println("You can create new credentials using `go-pass add --URL url_name --username user_name --password pass_word`")
			os.Exit(1)
		} else {
			fmt.Println(string(v))
		}
		return nil
	})
}

// GetAllURLUsers gets all the users for a given URL
func (s *Store) GetAllURLUsers(url string) error {
	return s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(url))
		if b == nil {
			fmt.Println("There are no saved credentials for this URL")
			fmt.Println("You can create new credentials using `go-pass add --URL url_name --username user_name --password pass_word`")
			os.Exit(1)
		}
		fmt.Println(url)
		if err := b.ForEach(func(k, _ []byte) error {
			fmt.Printf("|--%s\n", string(k))
			return nil
		}); err != nil {
			log.Fatal(err)
		}
		return nil
	})
}

// GetAllURLs gets all the sites that have credentials stored
func (s *Store) GetAllURLs() error {
	return s.db.View(func(tx *bolt.Tx) error {
		if err := tx.ForEach(func(k []byte, _ *bolt.Bucket) error {
			fmt.Println(string(k))
			return nil
		}); err != nil {
			log.Fatal(err)
		}
		return nil
	})
}
