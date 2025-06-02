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

func (db *Database) GetTable(name string) (*Table, error) {
	if _, ok := db.tables[name]; !ok {
		return nil, fmt.Errorf("table %s not found", name)
	}

	var err error
	var table *Table
	if table, err = loadTable(db.getTablePath(), name); err != nil {
		return nil, err
	}
	db.tables[name] = table
	return db.tables[name], nil
}

func (db *Database) CreateTable(name string, columns []*Column) error {
	if _, ok := db.tables[name]; ok {
		return fmt.Errorf("table %s already exists", name)
	}

	table, err := newTable(name, columns)
	if err != nil {
		return fmt.Errorf("table %s create error: %v", name, err)
	}

	db.tables[name] = table
	return db.saveToFile()
}

func loadDatabase(path, name string) (*Database, error) {
	database := &Database{
		Name:   name,
		tables: make(map[string]*Table),
		Path:   path,
	}
	if err := database.loadFromFile(); err != nil {
		return nil, err
	}
	return database, nil
}

func newDatabase(path, name string) (*Database, error) {
	database := &Database{
		Name:   name,
		tables: make(map[string]*Table),
		Path:   path,
	}
	return database, nil
}

func (db *Database) loadFromFile() error {
	metadata := struct {
		Tables map[string]*Table `json:"tables"`
	}{
		Tables: map[string]*Table{},
	}

	if !utils.FileExists(db.getFilePath()) {
		db.tables = metadata.Tables

		if err := utils.WriteJSON(db.getFilePath(), &metadata); err != nil {
			return fmt.Errorf("error creating database file: %v", err)
		}

		return nil
	}

	if err := utils.ReadJSON(db.getFilePath(), &metadata); err != nil {
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
	if err := utils.WriteJSON(db.getFilePath(), &metadata); err != nil {
		return fmt.Errorf("error writing database file: %v", err)
	}
	return nil
}

func (db *Database) getTablePath() string {
	return filepath.Join(db.Path, db.Name)
}

func (db *Database) getFilePath() string {
	return filepath.Join(db.Path, fmt.Sprintf("%s.json", db.Name))
}
