package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func dbDemo001() {

	db, err := sql.Open("mysql", "root:123mmm@tcp(192.168.99.100:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select name,age,email from user")

	if err != nil {
		log.Fatal(err)
	}
	var (
		name  string
		age   int
		email string
	)

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &age, &email)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name, age, email)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	dbDemo001()

}
