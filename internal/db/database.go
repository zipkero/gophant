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

func NewDatabase(name string) (*Database, error) {
	database := &Database{
		Name:     name,
		Tables:   []*Table{},
		Filepath: fmt.Sprintf("data/%s.json", name),
	}
	if err := database.SaveToFile(); err != nil {
		return nil, err
	}
	return database, nil
}

func (db *Database) NewTable(name string) {
	table := newTable(name)
	db.Tables = append(db.Tables, table)
}

func (db *Database) LoadFromFile() error {
	if err := utils.ReadJSON(db.Filepath, db); err != nil {
		return fmt.Errorf("error reading database file: %v", err)
	}
	return nil
}

func (db *Database) SaveToFile() error {
	if err := utils.WriteJSON(db.Filepath, db); err != nil {
		return fmt.Errorf("error writing database file: %v", err)
	}
	return nil
}
