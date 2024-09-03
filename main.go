package main

import (
	"comnet/utils/encryption"
	"fmt"
)

func main() {

	// Generate key pair
	publicKey, privateKey, _ := encryption.GenerateKeyPair()

	// Print public key
	fmt.Println("Public key:")
	fmt.Println(string(publicKey))

	// Print private key
	fmt.Println("Private key:")
	fmt.Println(string(privateKey))
}
