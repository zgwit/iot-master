package db

import (
	"github.com/timshannon/bolthold"
	"os"
)

type Options struct {
	Filename string
}

var store *bolthold.Store

func Store() *bolthold.Store {
	return store
}

func Open(opts Options) (err error) {
	store, err = bolthold.Open(opts.Filename, os.ModePerm, nil)
	return
}

func Close() error {
	return store.Close()
}
