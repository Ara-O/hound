package db

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

const BUCKET_NAME = "URLs"

func IndexExistsInDB(url string) bool {
	var v []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET_NAME))
		v = b.Get([]byte(url))
		return nil
	})

	if len(v) == 0 {
		return false
	} else {

		return true
	}
}

func AddIndex(url string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET_NAME))
		err := b.Put([]byte(url), []byte("exists"))

		if err != nil {
			fmt.Println("Error adding url to index")
			return err
		}

		return nil
	})

}

func init() {
	var err error

	// TODO: change to use other's home dir
	db, err = bolt.Open("/home/ara/go/bin/db.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BUCKET_NAME))

		if err != nil {
			return fmt.Errorf("create bucket error: %s", err)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func Close() {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
