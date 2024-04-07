package main

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	c := 0

	a := app.New()

	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	text1 := canvas.NewText("Hello Fyne!", color.White)
	text2 := canvas.NewText("This is a sample application!", color.White)
	content := container.NewVBox(text1, text2, l, widget.NewButton("Click me!", func() {
		c++
		l.SetText("count: " + strconv.Itoa(c))
	}))
	w.SetContent(content)

	w.ShowAndRun()
}
