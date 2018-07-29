package main

import (
	"image"
	"image/draw"

	"golang.org/x/exp/shiny/gesture"

	"golang.org/x/exp/shiny/widget/node"
	"golang.org/x/exp/shiny/widget/theme"
)

// ImageViewer is an interactive leaf widget that paints image.Image.
type ImageViewer struct {
	node.LeafEmbed

	Src image.Image
	SrcRect image.Rectangle

	Offset image.Point

	lastPos image.Point
}

// NewImage returns a new Image widget for the part of a source image defined
// by src and srcRect.
func NewImageViewer(src image.Image, srcRect image.Rectangle) *ImageViewer {
	w := &ImageViewer{
		Src: src,
		SrcRect: srcRect,
	}
	w.Wrapper = w
	return w
}

func (w *ImageViewer) OnInputEvent(e interface{}, origin image.Point) node.EventHandled {
	switch e := e.(type) {
	case gesture.Event:
		switch e.Type {
		case gesture.TypeIsDrag:
			w.lastPos = imagePoint(e.InitialPos)
			fallthrough
		case gesture.TypeDrag:
			p := imagePoint(e.CurrentPos)

			w.Offset = w.Offset.Add(p.Sub(w.lastPos))
			w.lastPos = p

			w.Wrapper.Mark(node.MarkNeedsPaintBase)
			return node.Handled
		}
	}
	return node.NotHandled
}

func (w *ImageViewer) Measure(t *theme.Theme, widthHint, heightHint int) {
	w.MeasuredSize = w.SrcRect.Size()
}

func (w *ImageViewer) PaintBase(ctx *node.PaintBaseContext, origin image.Point) error {
	w.Marks.UnmarkNeedsPaintBase()
	if w.Src == nil {
		return nil
	}
	draw.Draw(
		ctx.Dst, w.Rect.Add(origin),
		w.Src, w.SrcRect.Min.Sub(w.Offset),
		draw.Over)
	return nil
}

func imagePoint(p gesture.Point) image.Point {
	return image.Point{
		X: int(p.X),
		Y: int(p.Y),
	}
}
