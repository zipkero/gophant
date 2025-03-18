package db

import (
	"fmt"
	"gophant/pkg/utils"
)

type Storage struct {
	Filepath string
}

func newStorage(filename string) (*Storage, error) {
	storage := &Storage{
		Filepath: filename,
	}

	return storage, nil
}

func (s *Storage) SaveToFile() error {
	if err := utils.WriteJSON(s.Filepath, s); err != nil {
		return fmt.Errorf("error writing database file: %v", err)
	}
	return nil
}
