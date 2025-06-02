package db

import (
	"fmt"
	"gophant/pkg/utils"
	"path/filepath"
)

const (
	SchemaFileName = "schema.json"
	DataFileName   = "data.json"
)

type Table struct {
	Name       string
	schemaPath string
	dataPath   string
	schema     *Schema
	data       *Data
}

func loadTable(path, name string) (*Table, error) {
	table := &Table{
		schemaPath: filepath.Join(path, name, SchemaFileName),
		dataPath:   filepath.Join(path, name, DataFileName),
	}
	if err := table.loadFromFile(); err != nil {
		return nil, err
	}
	return table, nil
}

func newTable(name string, columns []*Column) (*Table, error) {
	table := &Table{
		Name:   name,
		schema: &Schema{Name: name, Columns: columns},
		data:   &Data{},
	}
	return table, nil
}

func (t *Table) loadSchema() error {
	metadata := struct {
		Schema *Schema `json:"schema"`
	}{
		Schema: &Schema{},
	}

	if !utils.FileExists(t.schemaPath) {
		t.schema = metadata.Schema

		if err := utils.WriteJSON(t.schemaPath, &metadata); err != nil {
			return fmt.Errorf("error creating table schema file: %v", err)
		}

		return nil
	}

	if err := utils.ReadJSON(t.schemaPath, &metadata); err != nil {
		return fmt.Errorf("error reading table schema file: %v", err)
	}

	t.schema = metadata.Schema

	return nil
}

func (t *Table) loadData() error {
	metadata := struct {
		Data *Data `json:"data"`
	}{
		Data: &Data{},
	}

	if !utils.FileExists(t.dataPath) {
		t.data = metadata.Data

		if err := utils.WriteJSON(t.dataPath, &metadata); err != nil {
			return fmt.Errorf("error creating table data file: %v", err)
		}

		return nil
	}

	if err := utils.ReadJSON(t.dataPath, &metadata); err != nil {
		return fmt.Errorf("error reading table data file: %v", err)
	}

	t.data = metadata.Data

	return nil
}

func (t *Table) loadFromFile() error {
	if err := t.loadSchema(); err != nil {
		return err
	}
	if err := t.loadData(); err != nil {
		return err
	}
	return nil
}

func (t *Table) saveToFile() error {
	return nil
}
