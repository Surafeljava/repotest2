package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

  var db *sql.DB

  func main() {
	  var err error
	
	  db, err = sql.Open("postgres", "host=localhost port=5433 user=postgres password=123456 dbname=restaurantdb sslmode=disable")
  
	  if err != nil {
		  panic(err)
	  }
  
	  defer db.Close()
  
	  if err = db.Ping(); err != nil {
		  panic(err)
	  }
  
	  role := "Default"
  
	  err = db.QueryRow("SELECT name FROM roles WHERE id = 9").Scan(&role)
	  if err != nil {
		  panic(err)
	  }
  
	  _, err = db.Exec("INSERT INTO roles (name) VALUES ('manager')")
  
	  if err != nil {
		  panic(err)
	  }
  
	  fmt.Println(role)
  }