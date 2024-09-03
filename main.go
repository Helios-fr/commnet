package main

import (
	"comnet/utils/encryption"
	"comnet/utils/users"
	"log"
)

func main() {
	if !encryption.Test_All() {
		log.Println("Encryption tests failed")
	}
	if !users.Test_All() {
		log.Println("Users tests failed")
	}
	log.Println("All tests passed")
}
