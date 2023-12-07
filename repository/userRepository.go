package repository

import (
	"consumer-producer/models"
	"database/sql"
	"log"

	_ "github.com/jmrobles/h2go"
	_ "github.com/mattn/go-sqlite3"
)

func AddUser(user *models.User) {
	connection := openDatabase()
	statement, error := connection.Prepare("INSERT INTO users(id, firstname, lastname, msisdn) VALUES(?,?,?,?)")
	if error != nil {
		log.Panicf("Failed to prepare statement: %s", error)
	}
	statementResult, error := statement.Exec(user.Id, user.Firstname, user.Lastname, user.Msisdn)
	connection.Close()
	if error != nil {
		log.Panicf("Failed to execute prepared statement: %s", error)
	}
	log.Printf("Success statement result: %s", statementResult)
}

func GetAllUsers() []models.User {
	connection := openDatabase()
	rows, error := connection.Query("SELECT * FROM users")
	connection.Close()
	if error != nil {
		log.Panicf("Failed to retrieve users: %s", error)
	}
	var users []models.User
	for rows.Next() {
		var user models.User
		error := rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Msisdn)
		if error != nil {
			log.Panicf("Failed to read fetched row: %s", error)
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
		log.Panicf("Failed to create table: %s", error)
	}
	log.Printf("Success create table: %s", result)
}

func openDatabase() *sql.DB {
	connection, error := sql.Open("sqlite3", "./users.db")
	if error != nil {
		log.Panicf("Failed to create connection: %s", error)
	}
	log.Printf("Success create connection: %s", connection)
	return connection
}
