package main

import (
	"context"

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

// Greet returns a greeting for the given name
func (a *App) Login(name string) (any, any) {
	// try to get the user from the database
	user_public, user_private := users.GetUser(name)
	if len(user_private) == 0 && len(user_public) != 0 {
		// user does not exist
		return nil, "User's private key is unknown"
	}
	if len(user_public) == 0 {
		// user does not exist
		return nil, "User does not exist, please create an account"
	}

	return name, nil
}
