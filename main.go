package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"commnet/backend/encryption"
	"commnet/backend/users"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	// Run the tests on the backend
	tests()

	// create the user default if it does not exist
	pub, priv := users.GetUser("default")
	if len(priv) == 0 && len(pub) == 0 {
		pub, priv, _ = encryption.GenerateKeyPair()
		users.CreateUser("default", pub, priv)
	}

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "commnet",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 100},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func tests() {
	if !encryption.Test_All() {
		log.Println("Encryption tests failed")
		log.Println("Generated keys:", encryption.Test_GenerateKeyPair())
		log.Println("Encrypted message:", encryption.Test_Encrypt())
		log.Println("Decrypted message:", encryption.Test_Decrypt())
	}
	if !users.Test_All() {
		log.Println("Users tests failed")
		log.Println("Create user:", users.Test_CreateUser())
		log.Println("Get user:", users.Test_GetUser())
		log.Println("Update user:", users.Test_UpdateUser())
		log.Println("Validate user:", users.Test_ValidateUser())
		log.Println("Remove user:", users.Test_RemoveUser())
		log.Println("Reset DB:", users.Test_ResetDB())
		log.Println("Get Authority:", users.Test_GetAuthority())
	}
}
