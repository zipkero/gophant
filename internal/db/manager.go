package db

import (
	"fmt"
	"gophant/pkg/utils"
)

const DatabaseFile = "data/db.json"

var currentDatabase *Database

type Manager struct {
	Databases map[string]*Database `json:"databases"`
}

func NewManager() (*Manager, error) {
	mgr := &Manager{
		Databases: map[string]*Database{},
	}
	err := mgr.LoadFromFile()
	if err != nil {
		return nil, err
	}

	return mgr, nil
}

func (mgr *Manager) LoadFromFile() error {
	if !utils.FileExists(DatabaseFile) {
		if err := utils.FileCreate(DatabaseFile); err != nil {
			return fmt.Errorf("error creating database file: %v", err)
		}
	}
	if err := utils.ReadJSON(DatabaseFile, mgr); err != nil {
		return fmt.Errorf("error reading database file: %v", err)
	}
	return nil
}

func (mgr *Manager) SaveToFile() error {
	if err := utils.WriteJSON(DatabaseFile, mgr); err != nil {
		return fmt.Errorf("error writing database file: %v", err)
	}
	return nil
}

func (mgr *Manager) UseDatabase(name string) error {
	if _, ok := mgr.Databases[name]; !ok {
		return fmt.Errorf("database %s not found", name)
	}
	currentDatabase = mgr.Databases[name]
	return nil
}

func (mgr *Manager) CreateDatabase(name string) error {
	if _, ok := mgr.Databases[name]; ok {
		return fmt.Errorf("database %s already exists", name)
	}
	db, err := NewDatabase(name)
	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}
	mgr.Databases[name] = db

	err = mgr.SaveToFile()
	if err != nil {
		return fmt.Errorf("error saving database file: %v", err)
	}
	return nil
}

func (mgr *Manager) CreateTable(name string, columns []*Column) error {
	if currentDatabase == nil {
		return fmt.Errorf("database %s not found", name)
	}
	if _, ok := currentDatabase.Tables[name]; ok {
		return fmt.Errorf("table %s already exists", name)
	}
	table, err := newTable(name, columns)
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}
	currentDatabase.Tables[name] = table
	return nil
}

func (mgr *Manager) CreateColumn(tableName string, columnName string, columnType ColumnType) error {
	if currentDatabase == nil {
		return fmt.Errorf("database %s not found", tableName)
	}
	table, ok := currentDatabase.Tables[tableName]
	if !ok {
		return fmt.Errorf("table %s not found", tableName)
	}
	if err := table.AddColumn(&Column{
		Name: columnName,
		Type: columnType,
	}); err != nil {
		return fmt.Errorf("error creating column: %v", err)
	}
	return nil
}
