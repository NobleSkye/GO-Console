package main

import (
	"your-module-name/internal/api" // Update with actual module name

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func mainWindow(client *api.APIClient) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Welcome to Go-Console"),
		widget.NewButton("Start Server", func() {
			// Logic to start server
		}),
		widget.NewButton("Stop Server", func() {
			// Logic to stop server
		}),
	)
}
