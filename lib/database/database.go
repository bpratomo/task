package database

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/boltdb/bolt"
	m "task/lib/models"
)

var db *bolt.DB

func init() {
	connect()
}

func connect() {
	newDb, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db = newDb
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("todo"))
		return err
	})
}

func GetNextId() int {
	var id int
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		idb, _ := b.NextSequence()
		id = int(idb)
		return nil
	})
	return id

}

func Update(id int, t m.Task) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))

		// Marshal user data into bytes.
		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put(itob(id), buf)
	})
}

func Get(filter string) []m.Task {
	var ts []m.Task
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("todo"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var t m.Task
			json.Unmarshal(v, &t)
			if t.ID == 0 {
				continue
			}

			if strings.Contains(t.Title, filter) || filter == "" {
				ts = append(ts, t)
			}
		}

		return nil
	})

	return ts
}

func Delete(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		err := b.Delete(itob(id))
		if err != nil {
			return err
		}
		return nil
	})

}
