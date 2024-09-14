package main

import (
	"context"
	"fmt"
	"net"
	"strings"

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
func (a *App) Verify(username string) any {
	fmt.Println("Verify", username)
	pub, priv := users.GetUser(username)
	if len(pub) == 0 || len(priv) == 0 {
		return "false"
	}

	return "true"
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
