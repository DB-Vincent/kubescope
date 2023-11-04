package applayout

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func Rgb(c uint32) color.NRGBA {
	return Argb((0xff << 24) | c)
}

func Argb(c uint32) color.NRGBA {
	return color.NRGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}

type Fill struct {
	Col color.NRGBA
}

func (f Fill) Layout(gtx layout.Context, size image.Point) layout.Dimensions {
	dr := image.Rectangle{
		Max: image.Point{X: size.X, Y: size.Y},
	}

	paint.FillShape(gtx.Ops, f.Col, clip.RRect{Rect: dr, SE: 10, SW: 10, NW: 10, NE: 10}.Op(gtx.Ops))
	return layout.Dimensions{Size: size}
}
