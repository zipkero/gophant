package db

import "fmt"

type Table struct {
	Name        string
	Columns     map[string]*Column
	ColumnOrder []string
	Storage     *Storage
}

func newTable(name string, columns []*Column) (*Table, error) {
	storage, err := newStorage(name)
	if err != nil {
		return nil, fmt.Errorf("error creating new table: %v", err)
	}

	colMap := make(map[string]*Column)
	colOrder := make([]string, len(columns))

	for i, col := range columns {
		colMap[col.Name] = col
		colOrder[i] = col.Name
	}

	return &Table{
		Name:        name,
		Columns:     colMap,
		ColumnOrder: colOrder,
		Storage:     storage,
	}, nil
}

func (t *Table) AddColumn(column *Column) error {
	if _, ok := t.Columns[column.Name]; ok {
		return fmt.Errorf("column %s already exists", column.Name)
	}
	t.Columns[column.Name] = column
	t.ColumnOrder = append(t.ColumnOrder, column.Name)
	return nil
}
