package main

import (
	"fmt"
	"gophant/internal/db"
)

func main() {
	mng, err := db.NewManager("data")
	if err != nil {
		panic(err)
	}

	var productDb *db.Database
	var orderDb *db.Database
	if productDb, err = mng.GetDatabase("product"); err != nil {
		fmt.Println(err)
	}
	if err := mng.CreateDatabase("product"); err != nil {
		fmt.Println(err)
	}
	if err := mng.CreateDatabase("order"); err != nil {
		fmt.Println(err)
	}

	productDb, err = mng.GetDatabase("product")
	if err != nil {
		fmt.Println(err)
	}

	t, err := productDb.GetTable("product_details")
	if err != nil {
		fmt.Println(err)
	}

	t, err = productDb.GetTable("product_details")
	if err != nil {
		fmt.Println(err)

		err = productDb.CreateTable("product_details", make([]*db.Column, 0))
		if err != nil {
			fmt.Println(err)
		}
	}

	t, err = productDb.GetTable("product_details")
	if err != nil {
		fmt.Println(err)
	}

	orderDb, err = mng.GetDatabase("order")
	if err != nil {
		fmt.Println(err)
	}

	t, err = orderDb.GetTable("order_details")
	if err != nil {
		fmt.Println(err)

		err = orderDb.CreateTable("order_details", make([]*db.Column, 0))
		if err != nil {
			fmt.Println(err)
		}
	}

	t, err = orderDb.GetTable("order_details")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t)
}
