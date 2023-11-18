package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/container/layout"
	"fyne.io/fyne/v2/container/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Simple GUI App")

	// Create a label
	myLabel := widget.NewLabel("Hello, Golang GUI!")

	// Create a button with a click handler
	myButton := widget.NewButton("Click me", func() {
		myLabel.SetText("Button Clicked!")
	})

	// Set up the GUI layout
	myWindow.SetContent(container.NewVBox(
		layout.NewSpacer(), // Spacer at the top
		myLabel,
		layout.NewSpacer(), // Spacer in the middle
		myButton,
		layout.NewSpacer(), // Spacer at the bottom
	))

	// Show and run the application
	myWindow.ShowAndRun()
}
