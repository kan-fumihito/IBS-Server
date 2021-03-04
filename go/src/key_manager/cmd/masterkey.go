package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(mysql:3306)/db?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	jst, _ := time.LoadLocation("Asia/Tokyo")

	var key string
	fmt.Scanf("%s", &key)

	_, err = db.Exec("update")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("insert")
	if err != nil {
		log.Fatal(err)
	}
}
