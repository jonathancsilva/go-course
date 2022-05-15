package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func connection() sql.DB {
	db, err := sql.Open("mysql", "admin:admin@/gocourse")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return *db
}

func handlerUser(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		getUserById(w, r, id)
	case r.Method == "GET":
		getAllUsers(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Sorry... :(")
	}
}

func getUserById(w http.ResponseWriter, r *http.Request, id int) {
	db := connection()
	// db, err := sql.Open("mysql", "admin:admin@/gocourse")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	var u User
	db.QueryRow("select id, name from users where id = ?", id).Scan(&u.Id, &u.Name)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	db := connection()
	// db, err := sql.Open("mysql", "admin:admin@/gocourse")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	rows, _ := db.Query("select id, name from users")
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name)
		users = append(users, user)
	}

	json, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
