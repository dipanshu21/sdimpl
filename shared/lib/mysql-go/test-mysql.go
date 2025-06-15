package main

import (
	"fmt"
	"mysql-go/lib"
)

func TestMySqlConnection() error {
	// Create a new connection to the database
	db, err := lib.CreateNewConnectionToDB()

	if err != nil {
		return fmt.Errorf("error creating connection to database: %w", err)
	} else {
		fmt.Println("Successfully connected to the database.")
	}
	defer db.Close()

	var users []lib.User = lib.GetUsers(db)
	for i, user := range users {
		fmt.Printf("Index: %d, ID: %s, Username: %s, Age: %d\n", i, user.Id, user.Username, user.Age)
	}
	// If we reach here, the connection was successful
	return nil
}

func main() {
	err := TestMySqlConnection()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
