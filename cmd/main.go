package main

import (
	"go-console/internal/api"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Initialize the Fyne application
	myApp := app.New()
	w := myApp.NewWindow("Go-Console")

	// Create client (assuming you have a function in your api package to create one)
	client := api.NewClient() // Replace with your actual client creation logic

	// Set the window content
	w.SetContent(mainWindow(client))
	w.ShowAndRun()
}

// mainWindow creates the UI layout for the main application window
func mainWindow(client *api.APIClient) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Welcome to Go-Console"),
		widget.NewButton("Start Server", func() {
			// Logic to start the server using the API client
			// e.g., client.StartServer()
		}),
		widget.NewButton("Stop Server", func() {
			// Logic to stop the server using the API client
			// e.g., client.StopServer()
		}),
		// Add more buttons and UI elements as needed
	)
}
