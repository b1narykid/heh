package main

import (
	"os"
	"log"
	"fmt"
	"image"
	"image/color"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/exp/shiny/widget"
	"golang.org/x/exp/shiny/widget/theme"

	_ "image/gif"
	_ "image/png"
	_ "image/jpeg"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

func main() {
	log.SetFlags(0)
	driver.Main(func(s screen.Screen) {
		if len(os.Args) < 2 {
			log.Fatal("no image file specified")
		}

		// TODO: view multiple images
		fname := os.Args[1]
		src, err := decode(fname)
		if err != nil {
			log.Fatal(err)
		}

		opts := &widget.RunWindowOptions{}
		opts.NewWindowOptions.Title = fname

		w := widget.NewSheet(
			widget.NewUniform(theme.StaticColor(color.Black),
				NewImageViewer(src, src.Bounds())))
		if err := widget.RunWindow(s, w, opts); err != nil {
			log.Fatal(err)
		}
	})
}

func decode(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	m, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("could not decode %s: %v", filename, err)
	}
	return m, nil
}
