package main

import (
	"log"
	"image/color"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/exp/shiny/widget"
	"golang.org/x/exp/shiny/widget/theme"
)

func main() {
	log.SetFlags(0)
	driver.Main(func(s screen.Screen) {
		opts := &widget.RunWindowOptions{}
		w := widget.NewSheet(
			widget.NewUniform(theme.StaticColor(color.Black), nil))
		if err := widget.RunWindow(s, w, opts); err != nil {
			log.Fatal(err)
		}
	})
}
