package users

/*
This modile handles all elements of user data, including the creation and replication of users from peers, the storage and retrival of user data, and the management of user data.

Functions included in this file:
- Create user (CreateUser: username*, publicKey*, privateKey) --> bool
- Get user (GetUser: username*) --> publicKey, privateKey
- Update user (UpdateUser: username*, publicKey*, privateKey) --> bool
- Validate user (ValidateUser: username*, publicKey) --> bool
- Remove user (RemoveUser: username*) --> bool
- Reset DB (resetDB: ) --> bool
- Get Authority (GetAuthority: username*) --> int
*/

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// UpdateUser --> bool
// This function updates the public key and private key for a given username.
func UpdateUser(username string, publicKey string, privateKey string) bool {
	// Open the user data csv file
	file, err := os.Open("user_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the user data csv file
	reader := csv.NewReader(file)
	reader.Comment = '#'

	// Create a new user data csv file to overwrite the existing file
	newFile, err := os.Create("user_data.csv.new")
	if err != nil {
		log.Println("Error creating file:", err)
		return false
	}
	defer newFile.Close()

	// Write the new user data to the new file
	writer := csv.NewWriter(newFile)
	writer.Comma = ','
	writer.UseCRLF = true
	defer writer.Flush()

	// Search for the user in the csv file
	for {
		// Read the next line of the csv file
		record, err := reader.Read()
		if err == nil {
			// Check if the username matches the current record
			if record[0] == username {
				// Update the public key and private key
				record[1] = publicKey
				record[2] = privateKey
			}

			// Write the record to the new file
			err := writer.Write(record)
			if err != nil {
				log.Println("Error writing record:", err)
				return false
			}
		} else {
			if err == io.EOF {
				break // End of file reached
			}
			log.Fatal(err)
		}
	}

	// Remove the old user data csv file
	err = os.Remove("user_data.csv")
	if err != nil {
		log.Println("Error removing file:", err)
		return false
	}

	// Rename the new user data csv file to the original file name
	err = os.Rename("user_data.csv.new", "user_data.csv")
	if err != nil {
		log.Println("Error renaming file:", err)
		return false
	}

	return true
}
