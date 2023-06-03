package controller

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

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

func dbCreate(t Task) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		id, _ := b.NextSequence()
		t.ID = int(id)

		// Marshal user data into bytes.
		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put(itob(t.ID), buf)
	})

}

func dbUpdate(id int, t Task) error {
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

func dbGet(filter string) []Task {
	var ts []Task
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("todo"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var t Task
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

func dbDelete(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		err := b.Delete(itob(id))
		if err != nil {
			return err
		}
		return nil
	})

}
