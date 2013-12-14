package main

import (
	"image"
	"image/color"
)

func isBoxPixel(x, y int, r image.Rectangle) bool {
	return x == r.Min.X || x == r.Min.X || x == r.Max.X || y == r.Min.Y || y == r.Max.Y ||
		x == r.Min.X+1 || x == r.Max.X-1 || y == r.Min.Y+1 || y == r.Max.Y-1
}

func box(img *image.RGBA, r image.Rectangle, c color.Color) {
	for dx := r.Min.X; dx <= r.Max.X; dx++ {
		for dy := r.Min.Y; dy <= r.Max.Y; dy++ {
			if isBoxPixel(dx, dy, r) {
				img.Set(dx, dy, c)
			}
		}
	}
}

func square(img *image.RGBA, p image.Point, s int, c color.Color) {
	for dx := p.X - s; dx < p.X+s; dx++ {
		for dy := p.Y - s; dy < p.Y+s; dy++ {
			img.Set(dx, dy, c)
		}
	}
}

func cross(img *image.RGBA, p image.Point, c color.Color) {
	crossByXY(img, p.X, p.Y, c)
}

func crossByXY(img *image.RGBA, x, y int, c color.Color) {
	img.Set(x-2, y, c)
	img.Set(x-1, y, c)
	img.Set(x, y, c)
	img.Set(x+1, y, c)
	img.Set(x+2, y, c)
	img.Set(x, y-3, c)
	img.Set(x, y-2, c)
	img.Set(x, y-1, c)
	img.Set(x, y, c)
	img.Set(x, y+1, c)
	img.Set(x, y+2, c)
	img.Set(x, y+3, c)
}
