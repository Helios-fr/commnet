package users

/*
This modile handles all elements of user data, including the creation and replication of users from peers, the storage and retrival of user data, and the management of user data.

Functions included in this file:
- Create user (CreateUser: username*, publicKey*, privateKey) --> bool
- Get user (GetUser: username*) --> publicKey, privateKey
- Validate user (ValidateUser: username*, publicKey) --> bool
- Remove user (RemoveUser: username*) --> bool
- Reset DB (resetDB: ) --> bool
*/

import (
	"encoding/csv"
	"log"
	"os"
)

// CreateUser --> bool
// This function appends a new user to the user data csv file, storing the username, public key, and private key if it is known.
func CreateUser(username string, publicKey string, privateKey string) bool {
	// Check if the private key is empty
	if privateKey == "" {
		privateKey = "N/A" // Handle empty private key appropriately
	}

	// Check if the user already exists
	if pub, priv := GetUser(username); pub != "" || priv != "" {
		return false
	}

	// Append the new user to the user data csv file
	file, err := os.OpenFile("user_data.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println("Error opening file:", err)
		return false
	}
	defer file.Close()

	// Write the new user to the csv file
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the new user to the csv file
	err = writer.Write([]string{username, publicKey, privateKey})
	if err != nil {
		log.Println("Error writing to file:", err)
		return false
	}

	return true
}
