package main

import "gophant/internal/db"

func main() {
	mng := db.NewManager()
	if err := mng.CreateDatabase("test"); err != nil {
		panic(err)
	}
}
