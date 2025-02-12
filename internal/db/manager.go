package db

import (
	"encoding/json"
	"os"
)

const DB_FILE = "data/db.json"

type Manager struct {
	Databases map[string]*Database
}

func NewManager() *Manager {
	mgr := &Manager{
		Databases: map[string]*Database{},
	}
	mgr.LoadFromFile()
	return mgr
}

func (mgr *Manager) LoadFromFile() {
	file, err := os.OpenFile(DB_FILE, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&mgr)
	if err != nil {
		panic(err)
	}
}

func (mgr *Manager) SaveToFile() {
}
