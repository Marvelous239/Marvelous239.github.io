package main

import (
	"syscall/js"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func jsAlert(w fyne.Window, text string) {
	js.Global().Call("alert", text)
}

func main() {
	a := app.New()
	w := a.NewWindow("Alert Test")

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter text...")

	button := widget.NewButton("Alert", func() {
		jsAlert(w, entry.Text)
	})

	w.SetContent(container.NewVBox(
		entry,
		button,
	))

	w.ShowAndRun()
}
