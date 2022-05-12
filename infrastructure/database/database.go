package database

import (
	"encoding/csv"
	"fmt"
	"go-bootcamp/config"
	"os"
)

type database struct {
	path string
}

type DB interface {
	Read() ([][]string, error)
	Write([]string) error
	WriteAll([][]string) error
}

func New() DB {
	return &database{path: config.Config.Database.File}
}

func (db *database) Read() ([][]string, error) {
	f, err := os.OpenFile(db.path, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return nil, fmt.Errorf("Opening file '%v' in DB: %w", db.path, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	return r.ReadAll()
}

func (db *database) Write(value []string) (err error) {
	f, err := os.OpenFile(db.path, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend)
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
	f, err := os.OpenFile(db.path, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend)
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
