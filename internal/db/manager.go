package db

import (
	"fmt"
	"gophant/pkg/utils"
)

const DatabaseFile = "data/db.json"

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
