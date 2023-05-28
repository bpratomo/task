package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/boltdb/bolt"
)

var latestTaskId = 0
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

func create(title string) error {
	t := Task{Title: title}
	return db.Update(func(tx *bolt.Tx) error {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket([]byte("todo"))

		// Generate ID for the user.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
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

func getAll(filter string) {
    var ts []Task
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("todo"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
            var t Task
            json.Unmarshal(v,&t)
            if strings.Contains(t.Title,filter) || filter == "" {
                ts = append(ts, t)
            }
		}

		return nil
	})

    fmt.Println(ts)
}

func Update(id int, t Task) {

}

func delete(id string) {

}
