package storage

import (
	"errors"
	"log"

	"github.com/cockroachdb/pebble"
)

type Store interface {
	Set(key []byte, value string) error
	Get(key []byte) (string, error)
}

type PebbleStore struct {
	dirName string
	db      *pebble.DB
}

func NewPebbleStore(dirName string) (*PebbleStore, error) {
	db, err := pebble.Open(dirName, &pebble.Options{})
	if err != nil {
		return nil, err
	}
	return &PebbleStore{
		dirName: dirName,
		db:      db,
	}, nil
}

func (ps *PebbleStore) Set(key, value []byte) error {

	if ps.db == nil {
		return errors.New("DB is not opened yet")
	}

	if err := ps.db.Set(key, value, nil); err != nil {
		return err
	}
	return nil
}

func (ps *PebbleStore) Get(key []byte) ([]byte, error) {

	if ps.db == nil {
		return nil, errors.New("DB is not opened yet")
	}

	value, closer, err := ps.db.Get(key)

	if err != nil {
		return nil, err
	}

	if err := closer.Close(); err != nil {
		log.Fatal(err)
	}

	return value, nil
}
