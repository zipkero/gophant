package db

type Table struct {
	Name    string
	columns []*Column
	rows    []*Row
}

type Row map[string]interface{}

func newTable(name string, columns []*Column) (*Table, error) {
	return &Table{
		Name:    name,
		columns: make([]*Column, 0),
		rows:    make([]*Row, 0),
	}, nil
}

func (t *Table) AddRow(row *Row) error {
	return nil
}

func (t *Table) loadSchema() error {
	return nil
}

func (t *Table) loadData() error {
	return nil
}

func (t *Table) loadFromFile() error {
	return nil
}

func (t *Table) saveToFile() error {
	return nil
}
