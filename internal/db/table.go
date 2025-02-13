package db

import "fmt"

type Table struct {
	Name    string
	Columns []*Column
	Storage *Storage
}

func newTable(name string, columns []*Column) (*Table, error) {
	storage, err := newStorage(name)
	if err != nil {
		return nil, fmt.Errorf("error creating new table: %v", err)
	}
	return &Table{
		Name:    name,
		Columns: columns,
		Storage: storage,
	}, nil
}

func (t *Table) AddColumn(column *Column) {
	t.Columns = append(t.Columns, column)
}
