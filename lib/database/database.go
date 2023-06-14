package database

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/boltdb/bolt"
	m "task/lib/models"
)


func connect() *bolt.DB {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("todo"))
		return err
	})
	return db
}

func GetNextId() int {
	var id int
	db := connect()
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		idb, _ := b.NextSequence()
		id = int(idb)
		return nil
	})

	db.Close()
	return id

}

func Update(id int, t m.Task) error {
	db := connect()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))

		// Marshal user data into bytes.
		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put(itob(id), buf)
	})
	db.Close()
	return err
}

func GetAll() ([]m.Task,map[m.Project]bool){
    return Get("")
}

func Get(filter string) ([]m.Task,map[m.Project]bool) {
	var ts []m.Task
    var ps map[m.Project]bool
    ps = make(map[m.Project]bool)
	db := connect()
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
                if len(t.Project.Name) > 0{
                    ps[t.Project] = true
                }
			}
		}

		return nil
	})
	db.Close()

	return ts, ps
}

func Delete(id int) error {
	db := connect()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		err := b.Delete(itob(id))
		if err != nil {
			return err
		}
		return nil
	})

	db.Close()
	return err

}
