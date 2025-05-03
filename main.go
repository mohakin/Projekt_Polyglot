package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()

	win := myApp.NewWindow("Polyglot")

	win.SetContent(widget.NewLabel("Hello World"))
	win.Resize(fyne.NewSize(800, 600))
	win.SetFixedSize(false)
	win.ShowAndRun()
}
