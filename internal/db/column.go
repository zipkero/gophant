package db

import "time"

type ColumnType string

const (
	ColumnTypeInt    ColumnType = "int"
	ColumnTypeString ColumnType = "string"
	ColumnTypeFloat  ColumnType = "float"
	ColumnTypeBool   ColumnType = "bool"
	ColumnTypeTime   ColumnType = "time"
)

type Column struct {
	Name string
	Type ColumnType
}

func (c *Column) Validate(value interface{}) bool {
	switch c.Type {
	case ColumnTypeInt:
		_, ok := value.(int)
		return ok
	case ColumnTypeString:
		_, ok := value.(string)
		return ok
	case ColumnTypeFloat:
		_, ok := value.(float64)
		return ok
	case ColumnTypeBool:
		_, ok := value.(bool)
		return ok
	case ColumnTypeTime:
		_, ok := value.(time.Time)
		return ok
	}
	return false
}
