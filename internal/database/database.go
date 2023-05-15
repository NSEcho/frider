package database

import (
	"encoding/json"
	"github.com/nsecho/frider/internal/helpers"
	"go.etcd.io/bbolt"
	"os"
)

const (
	databaseName = "frider.db"
)

type bucket []byte

var (
	scriptBucket     = bucket("scripts")
	databaseFullPath = ""
)

var buckets = []bucket{
	scriptBucket,
}

type DB struct {
	db *bbolt.DB
}

func NewDatabase() (*DB, error) {
	db, err := bbolt.Open(databaseFullPath, os.ModePerm, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			_, err := tx.CreateBucketIfNotExists(b)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return &DB{db}, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Save(sc Script) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(scriptBucket)
		encoded, err := json.Marshal(sc)
		if err != nil {
			return err
		}

		return b.Put([]byte(sc.Name), encoded)
	})
}

func (db *DB) Delete(name string) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(scriptBucket)
		return b.Delete([]byte(name))
	})
}

func (db *DB) Scripts() ([]Script, error) {
	var scripts []Script
	err := db.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(scriptBucket)
		cursor := b.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var script Script
			if err := json.Unmarshal(v, &script); err != nil {
				return err
			}
			scripts = append(scripts, script)
		}
		return nil
	})
	return scripts, err
}

func (db *DB) ScriptByName(name string) (*Script, bool, error) {
	var script Script
	found := false
	err := db.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(scriptBucket)
		cursor := b.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			found = false
			if err := json.Unmarshal(v, &script); err != nil {
				return err
			}
			if script.Name == name {
				found = true
				return nil
			}
		}
		return nil
	})

	return &script, found, err
}

func init() {
	dbPath, err := helpers.CreateDatabasePath(databaseName)
	if err != nil {
		panic("error creating database")
	}
	databaseFullPath = dbPath
}
