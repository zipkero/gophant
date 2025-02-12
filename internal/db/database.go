package db

import (
	"encoding/json"
	"os"
)

type Database struct {
	Name     string
	Tables   []*Table
	Filename string
}

func (db *Database) NewDatabase(name string) *Database {
	return &Database{
		Name:   name,
		Tables: []*Table{},
	}
}

func (db *Database) NewTable(name string) {
	table := newTable(name)
	db.Tables = append(db.Tables, table)
}

func (db *Database) LoadFromFile() {
	file, err := os.Open(db.Filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&db)
	if err != nil {
		panic(err)
	}
}

func (db *Database) SaveToFile() {
	file, err := os.Create(db.Filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(db)
	if err != nil {
		panic(err)
	}
}
