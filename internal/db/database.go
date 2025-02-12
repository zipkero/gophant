package db

import (
	"fmt"
	"gophant/pkg/utils"
)

type Database struct {
	Name     string   `json:"name"`
	Tables   []*Table `json:"tables"`
	Filepath string   `json:"filepath"`
}

func NewDatabase(name string) *Database {
	database := &Database{
		Name:     name,
		Tables:   []*Table{},
		Filepath: fmt.Sprintf("data/%s.json", name),
	}
	database.SaveToFile()
	return database
}

func (db *Database) NewTable(name string) {
	table := newTable(name)
	db.Tables = append(db.Tables, table)
}

func (db *Database) LoadFromFile() {
	if err := utils.ReadJSON(db.Filepath, db); err != nil {
		panic(err)
	}
}

func (db *Database) SaveToFile() {
	if err := utils.WriteJSON(db.Filepath, db); err != nil {
		panic(err)
	}
}
