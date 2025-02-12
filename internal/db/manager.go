package db

import (
	"fmt"
	"gophant/pkg/utils"
)

const DatabaseFile = "data/db.json"

type Manager struct {
	Databases map[string]*Database `json:"databases"`
}

func NewManager() *Manager {
	mgr := &Manager{
		Databases: map[string]*Database{},
	}
	mgr.LoadFromFile()
	return mgr
}

func (mgr *Manager) LoadFromFile() {
	if !utils.FileExists(DatabaseFile) {
		if err := utils.FileCreate(DatabaseFile); err != nil {
			panic(err)
		}
	}
	if err := utils.ReadJSON(DatabaseFile, mgr); err != nil {
		panic(err)
	}
}

func (mgr *Manager) SaveToFile() {
}

func (mgr *Manager) CreateDatabase(name string) error {
	if _, ok := mgr.Databases[name]; ok {
		return fmt.Errorf("database %s already exists", name)
	}
	mgr.Databases[name] = NewDatabase(name)
	mgr.SaveToFile()
	return nil
}
