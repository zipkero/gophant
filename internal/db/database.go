package db

import (
	"fmt"
	"gophant/pkg/utils"
)

type Database struct {
	Name     string            `json:"name"`
	Tables   map[string]*Table `json:"tables"`
	Filepath string            `json:"filepath"`
}

func NewDatabase(name string) (*Database, error) {
	database := &Database{
		Name:     name,
		Tables:   make(map[string]*Table),
		Filepath: fmt.Sprintf("data/%s.json", name),
	}
	if err := database.SaveToFile(); err != nil {
		return nil, err
	}
	return database, nil
}

func (db *Database) NewTable(name string, columns []*Column) error {
	if _, ok := db.Tables[name]; ok {
		return fmt.Errorf("table %s already exists", name)
	}
	table, err := newTable(name, columns)
	if err != nil {
		return fmt.Errorf("table %s create error: %v", name, err)
	}
	db.Tables[name] = table
	if err := db.SaveToFile(); err != nil {
		return fmt.Errorf("failed to save table %s to file: %s", name, err)
	}
	return nil
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
