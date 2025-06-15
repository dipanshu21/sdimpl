// write code to generate mock user data for testing purposes
// given the number of users to generate
package users

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

// generateMockUserData generates a slice of mock user data for testing purposes
func generateMockUserData(numUsers int) []User {
	mockUsers := make([]User, numUsers)

	for i := 0; i < numUsers; i++ {
		mockUsers[i] = User{
			Id:       "NA",                                                 // Random ID for testing
			Username: fmt.Sprintf("user_%d_%d", time.Now().UnixMilli(), i), // Unique username with timestamp
			Age:      int8(rand.Intn(100)),                                 // Random age between 0 and 99
		}
	}

	return mockUsers
}

func CreateMockUsers(db *sql.DB, numUsers int) error {
	mockUsers := generateMockUserData(numUsers)

	for _, user := range mockUsers {
		err := CreateNewUser(db, user.Username, user.Age)
		if err != nil {
			return err
		}
	}

	return nil
}
