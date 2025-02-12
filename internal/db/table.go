package db

type Table struct {
	Name    string
	Columns []*Column
	Storage *Storage
}

func newTable(name string) *Table {
	storage := newStorage(name)

	return &Table{
		Name:    name,
		Columns: []*Column{},
		Storage: storage,
	}
}

func (t *Table) AddColumn(column *Column) {
	t.Columns = append(t.Columns, column)
}
