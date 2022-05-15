package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "admin:admin@/gocourse")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// update
	stmt, _ := db.Prepare("update users set name = ? where id = ?")
	stmt.Exec("Tiago", 1)
	stmt.Exec("Ricardo", 3)

	// delete
	stmt2, _ := db.Prepare("delete from users where id = ?")
	stmt2.Exec(2)
}
