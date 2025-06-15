package users

import (
	"database/sql"
	"fmt"
	"log"
	"mysql-go/lib"
)

type User struct {
	Id       string
	Username string
	Age      int8
}

func GetUsers(db *sql.DB) []User {
	query := "SELECT idUsers, username, age FROM users LIMIT 1000"
	rows, err := lib.ExecuteMySqlQuerySELECT(db, query)

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}

	defer rows.Close()

	var users []User = make([]User, 0)

	for rows.Next() {
		var id int
		var username string
		var age int

		err := rows.Scan(&id, &username, &age)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, User{
			Id:       fmt.Sprintf("%d", id),
			Username: username,
			Age:      int8(age),
		})
	}

	return users
}

func CreateNewUser(db *sql.DB, username string, age int8) error {
	query := "INSERT INTO users (username, age) VALUES (?, ?)"
	err := lib.ExecuteMySqlQueryINSERTOrDELETE(db, query, username, age)

	if err != nil {
		return fmt.Errorf("error inserting new user: %w", err)
	}

	return nil
}

func DeleteAllUsers(db *sql.DB) error {
	query := "DELETE FROM users"
	err := lib.ExecuteMySqlQueryINSERTOrDELETE(db, query)

	if err != nil {
		return fmt.Errorf("error deleting all users: %w", err)
	}

	return nil
}
