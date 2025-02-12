package main

import "gophant/internal/db"

func main() {
	mng, err := db.NewManager()
	if err != nil {
		panic(err)
	}
	if err := mng.CreateDatabase("test"); err != nil {
		panic(err)
	}
}
