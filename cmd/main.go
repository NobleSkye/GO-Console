package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new application
	myApp := app.New()
	myWindow := myApp.NewWindow("Minecraft Server Manager")

	// Button to start the server
	startButton := widget.NewButton("Start Server", func() {
		// Logic to send start command through WebSocket
		fmt.Println("Starting server...")
		// Call your function to send the start command here
	})

	// Button to stop the server
	stopButton := widget.NewButton("Stop Server", func() {
		// Logic to send stop command through WebSocket
		fmt.Println("Stopping server...")
		// Call your function to send the stop command here
	})

	// Console button (real-time updates)
	consoleButton := widget.NewButton("Open Console", func() {
		// Logic to connect to WebSocket and display console output
		fmt.Println("Opening console...")
		// Set up a separate UI window or panel for console output
	})

	// Arrange buttons in a vertical layout
	myWindow.SetContent(container.NewVBox(
		startButton,
		stopButton,
		consoleButton,
	))

	// Launch the app
	myWindow.ShowAndRun()
}
