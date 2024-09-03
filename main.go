package main

import (
	"comnet/utils/encryption"
	"comnet/utils/server"
	"comnet/utils/users"
	"log"
)

// global variables
var (
	username    string = "nyx"
	server_port string = "6875"
)

func main() {
	if !encryption.Test_All() {
		log.Println("Encryption tests failed")
	}
	if !users.Test_All() {
		log.Println("Users tests failed")
	}
	log.Println("All tests passed")

	// verify the user data is correct
	log.Println("Username:", username)
	self_publickey, self_privateKey := users.GetUser(username)
	if self_privateKey == "" && self_publickey == "" {
		pubKey, privKey, _ := encryption.GenerateKeyPair()
		self_publickey = string(pubKey)
		self_privateKey = string(privKey)
		users.CreateUser(username, self_publickey, self_privateKey)
		log.Println("User created")
	} else if self_privateKey == "" && self_publickey != "" {
		// exit the program if the user does not have a private key
		log.Println("User does not have a private key for specified username")
		return
	}

	// start listening for connections on port 6875
	go server.Listen(server_port, self_privateKey)
}
