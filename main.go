package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

// //go:embed all:frontend/src/assets
// var cytoss embed.FS

// type FileLoader struct {
// 	http.Handler
// }

// func NewFileLoader() *FileLoader {
// 	return &FileLoader{}
// }

// func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	// var err error
// 	requestedFilename := "frontend/src/assets/" + strings.TrimPrefix(req.URL.Path, "/")
// 	println("Requesting file:", requestedFilename)

// 	if ctnt, err := cytoss.ReadFile(requestedFilename); err != nil {
// 		res.WriteHeader(http.StatusBadRequest)
// 		res.Write([]byte(fmt.Sprintf("Could not load file %s", err.Error())))
// 	} else {
// 		res.Write(ctnt)
// 	}

// 	// fileData, err := os.ReadFile(requestedFilename)
// 	// if err != nil {
// 	// 	res.WriteHeader(http.StatusBadRequest)
// 	// 	res.Write([]byte(fmt.Sprintf("Could not load file %s", requestedFilename)))
// 	// }

// }

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Пререквизитләр",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
			// Handler: NewFileLoader(),
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
