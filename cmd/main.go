package main

import (
	"fmt"
	"gophant/internal/db"
)

func main() {
	mng, err := db.NewManager()
	if err != nil {
		panic(err)
	}
	if err = mng.UseDatabase("test"); err != nil {
		fmt.Println(err)
	}
	if err := mng.CreateDatabase("test"); err != nil {
		fmt.Println(err)
	}
	if err = mng.UseDatabase("test"); err != nil {
		fmt.Println(err)
	}

	if err = mng.CreateTable("test", []*db.Column{
		{
			Name: "id",
			Type: db.ColumnTypeInt,
		},
		{
			Name: "name",
			Type: db.ColumnTypeString,
		},
	},
	); err != nil {
		fmt.Println(err)
	}
}
