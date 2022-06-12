package database

import (
	"encoding/csv"
	"fmt"
	"os"
)

type database struct {
	path string
}

type DB interface {
	Read() ([][]string, error)
	ConcurrentRead(
		filter string, maxItems, itemsPerWorker int) ([][]string, error)
	Write([]string) error
	WriteAll([][]string) error
}

func New(path string) DB {
	return &database{path: path}
}

func (db *database) Read() ([][]string, error) {
	f, err := os.OpenFile(db.path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("opening file '%v' in db: %w", db.path, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	return r.ReadAll()
}

func (db *database) Write(value []string) (err error) {
	f, err := os.OpenFile(db.path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer func() {
		cerr := f.Close()
		if err == nil {
			err = cerr
		}
	}()

	w := csv.NewWriter(f)
	err = w.Write(value)
	if err != nil {
		return
	}
	w.Flush()

	return
}

func (db *database) WriteAll(value [][]string) (err error) {
	f, err := os.OpenFile(db.path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer func() {
		cerr := f.Close()
		if err == nil {
			err = cerr
		}
	}()

	w := csv.NewWriter(f)
	err = w.WriteAll(value)
	if err != nil {
		return
	}
	w.Flush()

	return
}
