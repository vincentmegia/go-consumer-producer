package repository

import (
	"database/sql"
	"log"

	"github.com/vincentmegia/consumer-producer/models"

	_ "github.com/jmrobles/h2go"
	_ "github.com/mattn/go-sqlite3"
)

func AddUser(user *models.User) {
	connection := openDatabase()
	statement, error := connection.Prepare("INSERT INTO users(id, firstname, lastname, msisdn) VALUES(?,?,?,?)")
	if error != nil {
		log.Println("Failed to prepare statement: ", error)
	}
	statement.Exec(user.Id, user.Firstname, user.Lastname, user.Msisdn)
	connection.Close()
	// if error != nil {
	// 	log.Println("Failed to execute prepared statement: ", error)
	// }
	// log.Println("Success statement result: ", statementResult)
}

func GetAllUsers() []models.User {
	connection := openDatabase()
	rows, error := connection.Query("SELECT * FROM users")
	connection.Close()
	if error != nil {
		log.Println("Failed to retrieve users: ", error)
	}
	var users []models.User
	for rows.Next() {
		var user models.User
		error := rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Msisdn)
		if error != nil {
			log.Println("Failed to read fetched row: ", error)
		}
		users = append(users, user)
	}
	return users
}

func CreateTable() {
	connection := openDatabase()
	result, error := connection.Exec(
		"DROP TABLE users; CREATE TABLE users(ID VARCHAR(255), FIRSTNAME VARCHAR(255), LASTNAME VARCHAR(255), MSISDN VARCHAR(255))",
	)
	connection.Close()
	if error != nil {
		log.Println("Failed to create table: ", error)
	}
	log.Println("Success create table: ", result)
}

func openDatabase() *sql.DB {
	connection, error := sql.Open("sqlite3", "./users.db")
	if error != nil {
		log.Println("Failed to create connection: ", error)
	}
	return connection
}
