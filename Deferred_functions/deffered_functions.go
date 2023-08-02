package main

import "database/sql"

func main() {
	// fmt.Println("main 1")
	// defer fmt.Println("defer 1")
	// fmt.Println("main 2")
	// defer fmt.Println("defer 2")

	// placeing the defer right after allows you to know that your resources will be released in the correct order. first in, last out.
	// .Close() is to release your resource.
	db, _ := sql.Open("driverName", "connection string")
	defer db.Close()

	rows, _ := db.Query("some sql query here")
	defer rows.Close()

}