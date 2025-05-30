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
	var database *db.Database
	if database, err = mng.GetDatabase("product"); err != nil {
		fmt.Println(err)
	}
	if err := mng.CreateDatabase("product"); err != nil {
		fmt.Println(err)
	}
	if err := mng.CreateDatabase("order"); err != nil {
		fmt.Println(err)
	}
	database, err = mng.GetDatabase("product")
	if err != nil {
		fmt.Println(err)
	}
	t, err := database.GetTable("product_details")
	if err != nil {
		fmt.Println(err)
	}
	database, err = mng.GetDatabase("order")
	if err != nil {
		fmt.Println(err)
	}
	t, err = database.GetTable("order_details")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
}
