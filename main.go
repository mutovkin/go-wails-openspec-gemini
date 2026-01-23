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
		Title:  "6/49 Lottery Picker",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Fatal("Error:", err.Error())
	}
}