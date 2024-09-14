package main

import (
	"context"
	"fmt"
	"net"
	"strings"

	"commnet/backend/encryption"
	"commnet/backend/users"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Hello world function
func (a *App) Hello() string {
	return "Hello 👋"
}

// Login is called when the user logs in, it returns the user name only if the user exists and the private key is known
func (a *App) Login(name string) any {
	fmt.Println("Login: ", name)
	// try to get the user from the database
	user_public, user_private := users.GetUser(name)
	if len(user_private) == 0 && len(user_public) != 0 {
		// user does not exist
		fmt.Println("User does not exist, please create an account")
		return "Error: User does not exist, please create an account"
	}
	if len(user_public) == 0 {
		// user does not exist
		fmt.Println("User does not exist, please create an account")
		return "Error: User does not exist, please create an account"
	}

	return name
}

// Register is called when the user registers, it returns the user name only if the user does not exist
func (a *App) Register(name string) any {
	fmt.Println("Register: ", name)
	// try to get the user from the database
	user_public, _ := users.GetUser(name)
	if len(user_public) != 0 {
		// user exists
		fmt.Println("User already exists, please login")
		return "Error: User already exists, please login"
	}

	// generate keys for the user
	user_public, user_private, _ := encryption.GenerateKeyPair()
	if users.CreateUser(name, user_public, user_private) {
		// user created
		fmt.Println("User created")
		return name
	} else {
		// user not created
		fmt.Println("User not created")
		return "Error: User not created"
	}
}

func (a *App) DefaultUsername() string {
	var ips []net.IP
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return "unknown"
	}

	for _, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP)
			}
		}
	}

	for _, ip := range ips {
		if strings.HasPrefix(ip.String(), "192.168") {
			return ip.String()
		}
	}

	return "unknown"
}
