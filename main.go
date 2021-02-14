package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	var option int

	fmt.Println("--------------------------------")
	fmt.Println("CONTACT BOOK")
	fmt.Println("--------------------------------")
	fmt.Println("1. Create Contact")
	fmt.Println("2. View Contact's")
	fmt.Println("3. Update Contact")
	fmt.Println("4. Delete Contact")
	fmt.Println("5. Exit")
	fmt.Printf("\nEnter your Option: ")
	fmt.Scanf("%d", &option)
	switch option {
	case 1:

	}
}

func dbCheck() {
	var db, err = sql.Open("postgres", "host="+"localhost"+" user="+"postgres"+" dbname="+"contactbook"+" sslmode=disable password="+"sumeet")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
