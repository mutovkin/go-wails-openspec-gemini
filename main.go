// Package main is the entry point for the Wails application.
package main

import (
	"embed"
	"log"
	"lottery-picker/internal/adapters/primary/wailsapp"
	"lottery-picker/internal/adapters/secondary/random"
	"lottery-picker/internal/core/services"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

const (
	appTitle  = "6/49 Lottery Picker"
	appWidth  = 1024
	appHeight = 768
	bgRed     = 255
	bgGreen   = 255
	bgBlue    = 255
	bgAlpha   = 1
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// 1. Create Secondary Adapter (Output Port)
	rngAdapter := random.NewCryptoRandomAdapter()

	// 2. Create Service (Input Port implementation)
	lotteryService := services.NewLotteryService(rngAdapter)

	// 3. Create Primary Adapter (Wails App)
	app := wailsapp.NewApp(lotteryService)

	// 4. Run Wails Application
	err := wails.Run(&options.App{
		Title:  appTitle,
		Width:  appWidth,
		Height: appHeight,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: bgRed, G: bgGreen, B: bgBlue, A: bgAlpha},
		OnStartup:        app.Startup,
		Bind:             []any{app},
	})
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
}
