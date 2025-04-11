package main

import (
	"log"

	"github.com/AdityaByte/mysql/config"
	"github.com/AdityaByte/mysql/database"
	"github.com/AdityaByte/mysql/model"
	"github.com/AdityaByte/mysql/service"
)

const (
	dbName = "adityadb"
	tableName = "users"
)

func main() {
	db, err := config.CreateConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Connected to db successfully")

	if !database.TableExists(db, dbName, tableName) {
		_, err := db.Exec("CREATE TABLE users(id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, email varchar(50), password varchar(20))")
		if err != nil {
			log.Fatal("Failed to insert into db:", err)
		}
		log.Println("Table created successfully")
	} else {
		log.Println("Table already existed")
	}

	user := model.User{
		Email: "aditya@gmail.com",
		Password: "aditya",
	}

	if err := service.AddUser(&user, db, tableName); err != nil {
		log.Fatal(err)
	}

}
