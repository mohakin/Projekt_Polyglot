package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type Application struct {
	App *fyne.App
	Win fyne.Window
}

func main() {
	myApp := app.New()

	win := myApp.NewWindow("Polyglot")

	win.SetContent(widget.NewLabel("Hello World"))
	win.SetFullScreen(true)
	win.SetFixedSize(false)
	win.ShowAndRun()
}
