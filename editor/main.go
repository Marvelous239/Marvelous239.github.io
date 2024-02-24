package main

import (
	"syscall/js"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"

	"fyne.io/fyne/v2/widget"
)

var (
	MainEntry = widget.NewMultiLineEntry()
	FName     = "PLACEHOLDER"
)

func loadFile() {
	input := js.Global().Get("document").Call("createElement", "input")
	input.Set("type", "file")
	input.Call("addEventListener", "change", js.FuncOf(readFile), "false")
	input.Call("click")
}

func readFile(this js.Value, args []js.Value) any {
	file := this.Get("files").Index(0)
	if file.IsNull() {
		return nil
	}
	FName = file.Get("name").String()

	reader := js.Global().Get("FileReader").New()
	reader.Set("onload", js.FuncOf(func(this js.Value, args []js.Value) any {
		MainEntry.SetText(this.Get("result").String())
		MainEntry.Refresh()
		return nil
	}))
	reader.Call("readAsText", file)

	return nil
}

func saveFile() {
	encodeURIComponent := js.Global().Get("encodeURIComponent")

	link := js.Global().Get("document").Call("createElement", "a")
	data := "data:text/plain;charset=utf-8," + encodeURIComponent.Invoke(MainEntry.Text).String()
	link.Set("href", data)
	link.Set("download", FName)
	FName = "PLACEHOLDER"

	MainEntry.SetText("")
	MainEntry.Refresh()

	link.Call("click")
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Editor")

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.FileIcon(), loadFile),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), saveFile),
	)

	MainEntry.TextStyle.Monospace = true
	scrollContainer := container.NewScroll(MainEntry)
	scrollContainer.SetMinSize(fyne.NewSize(200, 400))

	content := container.NewVBox(
		layout.NewSpacer(),
		toolbar,
		scrollContainer,
		layout.NewSpacer(),
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
