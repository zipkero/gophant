package db

import (
	"fmt"
	"gophant/pkg/utils"
	"path/filepath"
)

type Database struct {
	Name   string
	Path   string
	tables map[string]*Table
}

func newDatabase(path, name string) (*Database, error) {
	database := &Database{
		Name:   name,
		tables: make(map[string]*Table),
		Path:   filepath.Join(path, fmt.Sprintf("%s.json", name)),
	}
	if err := database.loadFromFile(); err != nil {
		return nil, err
	}
	return database, nil
}

func (db *Database) GetTable(name string) (*Table, error) {
	if _, ok := db.tables[name]; !ok {
		return nil, fmt.Errorf("table %s not found", name)
	}
	return db.tables[name], nil
}

func (db *Database) NewTable(name string, columns []*Column) error {
	if _, ok := db.tables[name]; ok {
		return fmt.Errorf("table %s already exists", name)
	}
	table, err := newTable(name, columns)
	if err != nil {
		return fmt.Errorf("table %s create error: %v", name, err)
	}
	db.tables[name] = table
	if err := db.saveToFile(); err != nil {
		return fmt.Errorf("failed to save table %s to file: %s", name, err)
	}
	return nil
}

func (db *Database) loadFromFile() error {
	metadata := struct {
		Tables map[string]*Table `json:"tables"`
	}{
		Tables: map[string]*Table{},
	}

	if !utils.FileExists(db.Path) {
		db.tables = metadata.Tables

		if err := utils.WriteJSON(db.Path, &metadata); err != nil {
			return fmt.Errorf("error creating database file: %v", err)
		}

		return nil
	}

	if err := utils.ReadJSON(db.Path, &metadata); err != nil {
		return fmt.Errorf("error reading database file: %v", err)
	}

	db.tables = metadata.Tables

	return nil
}

func (db *Database) saveToFile() error {
	metadata := struct {
		Tables map[string]*Table `json:"tables"`
	}{
		Tables: db.tables,
	}
	if err := utils.WriteJSON(db.Path, &metadata); err != nil {
		return fmt.Errorf("error writing database file: %v", err)
	}
	return nil
}
