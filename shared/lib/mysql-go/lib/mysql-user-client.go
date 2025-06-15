package lib

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	Id       string
	Username string
	Age      int8
}

func GetUsers(db *sql.DB) []User {
	query := "SELECT idUsers, username, age FROM users LIMIT 1000"
	rows, err := ExecuteMySqlQuery(db, query)

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
