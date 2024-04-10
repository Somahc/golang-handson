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
	// c := 0

	a := app.New()

	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	e := widget.NewEntry()
	e.SetText("0")
	c := widget.NewCheck("Check!", func(f bool) {
		if f {
			l.SetText("CHECKED!")
		} else {
			l.SetText("UNCHECKED!")
		}
	})
	c.SetChecked(true)
	text1 := canvas.NewText("Hello Fyne!", color.White)
	text2 := canvas.NewText("This is a sample application!", color.White)
	content := container.NewVBox(text1, text2, l, e, c,
		widget.NewButton("Click me!", func() {
			n, _ := strconv.Atoi(e.Text)
			l.SetText("Total: " + strconv.Itoa(total(n)))
			// c++
			// l.SetText("count: " + strconv.Itoa(c))
		}))
	w.SetContent(content)

	// a.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantDark})
	w.ShowAndRun()
}

func total(n int) int {
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	return t
}
