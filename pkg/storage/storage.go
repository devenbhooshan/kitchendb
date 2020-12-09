package storage

import (
	"errors"
	"log"

	"github.com/cockroachdb/pebble"
)

type Store interface {
	Set(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	GetWithPrefix(prefix []byte) ([][]byte, error)
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

func (ps *PebbleStore) GetWithPrefix(prefix []byte) ([][]byte, error) {
	keyUpperBound := func(b []byte) []byte {
		end := make([]byte, len(b))
		copy(end, b)
		for i := len(end) - 1; i >= 0; i-- {
			end[i] = end[i] + 1
			if end[i] != 0 {
				return end[:i+1]
			}
		}
		return nil // no upper-bound
	}

	prefixIterOptions := func(prefix []byte) *pebble.IterOptions {
		return &pebble.IterOptions{
			LowerBound: prefix,
			UpperBound: keyUpperBound(prefix),
		}
	}

	iter := ps.db.NewIter(prefixIterOptions(prefix))
	data := [][]byte{}
	for iter.First(); iter.Valid(); iter.Next() {
		data = append(data, iter.Value())
	}
	return data, nil
}
