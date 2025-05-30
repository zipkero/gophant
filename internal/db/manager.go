package db

import (
	"fmt"
	"gophant/pkg/utils"
	"path/filepath"
)

const (
	DatabasesFileName   = "gophant.json"
	DatabasesFolderName = "databases"
)

type Manager struct {
	databases map[string]*Database
	path      string
	name      string
}

func NewManager(path string) (*Manager, error) {
	mgr := &Manager{
		databases: map[string]*Database{},
		path:      path,
	}
	err := mgr.loadFromFile()
	if err != nil {
		return nil, err
	}

	return mgr, nil
}

func (mgr *Manager) CreateDatabase(name string) error {
	if _, ok := mgr.databases[name]; ok {
		return fmt.Errorf("database %s already exists", name)
	}
	db, err := newDatabase(filepath.Join(mgr.path, DatabasesFolderName), name)
	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}

	err = mgr.addDatabase(name, db)

	if err != nil {
		return fmt.Errorf("error saving database file: %v", err)
	}
	return nil
}

func (mgr *Manager) GetDatabase(name string) (*Database, error) {
	if _, ok := mgr.databases[name]; !ok {
		return nil, fmt.Errorf("database %s not found", name)
	}
	return mgr.databases[name], nil
}

func (mgr *Manager) DropDatabase(name string) error {
	if _, ok := mgr.databases[name]; !ok {
		return fmt.Errorf("database %s not found", name)
	}
	delete(mgr.databases, name)
	return nil
}

func (mgr *Manager) loadFromFile() error {
	metadata := struct {
		Databases map[string]*Database `json:"databases"`
	}{
		Databases: map[string]*Database{},
	}

	if !utils.FileExists(mgr.path) {
		mgr.databases = metadata.Databases

		if err := utils.WriteJSON(mgr.getDatabaseFilePath(), &metadata); err != nil {
			return fmt.Errorf("error creating database file: %v", err)
		}

		return nil
	}

	if err := utils.ReadJSON(mgr.getDatabaseFilePath(), &metadata); err != nil {
		return fmt.Errorf("error reading database file: %v", err)
	}

	mgr.databases = metadata.Databases

	return nil
}

func (mgr *Manager) saveToFile() error {
	metadata := struct {
		Databases map[string]*Database `json:"databases"`
	}{
		Databases: mgr.databases,
	}

	if err := utils.WriteJSON(mgr.getDatabaseFilePath(), metadata); err != nil {
		return fmt.Errorf("error writing database file: %v", err)
	}

	return nil
}

func (mgr *Manager) addDatabase(name string, db *Database) error {
	mgr.databases[name] = db
	return mgr.saveToFile()
}

func (mgr *Manager) getDatabaseFilePath() string {
	return filepath.Join(mgr.path, DatabasesFileName)
}
