package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
)

var db *sql.DB
var err error
var count = 0

func main() {
	db, err = sql.Open("mysql", "root:kirikanan@tcp(localhost:3306)/example?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)
	http.HandleFunc("/", index)
	http.Handle("favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	q := `SELECT * FROM User;`
	x, err := db.Query(q)
	check(err)
	defer x.Close()
	log.Println(count)
	count++
	for x.Next() {
		var (
			id int64
			name string
			age int64
		)
		if err := x.Scan(&id, &name, &age); err != nil{
			log.Fatalln(err)
		}
		//fmt.Fprintln(w, "RETRIEVED RECORD:", name)
		log.Printf("id %d name %s age %d\n", id, name, age)
	}
	//if !x.NextResultSet() {
	//	log.Fatalf("expected more result sets: %v", x.Err())
	//}
	_, err = io.WriteString(w, "Successfully completed.")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}