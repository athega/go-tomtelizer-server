package main

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

type Feature struct {
	LeftEye  image.Point
	RightEye image.Point
	Mouth    image.Point
}

func handleFacialFeatures(features []Feature, original image.Image, timestamp string) {
	w := original.Bounds().Max.X
	h := original.Bounds().Max.Y

	puts("Image width:", w, "height:", h)

	m := image.NewRGBA(original.Bounds())
	draw.Draw(m, original.Bounds(), original, image.ZP, draw.Src)

	if len(features) > 0 {
		puts("Found features:", len(features), features)

		for _, f := range features {
			hatRect := calculateHatRect(f)

			if DEBUG {
				red := color.RGBA{255, 0, 0, 255}
				green := color.RGBA{0, 255, 0, 255}
				blue := color.RGBA{0, 0, 255, 255}

				box(m, hatRect, color.White)

				square(m, f.LeftEye, 4, red)
				square(m, f.Mouth, 4, green)
				square(m, f.RightEye, 4, blue)
			} else {
				// Load the santa hat
				santa := resizeImage(SANTA_HAT, hatRect.Size().X, hatRect.Size().Y)
				draw.Draw(m, hatRect, santa, image.ZP, draw.Over)
			}
		}
	} else {
		puts("No features found…")
	}

	saveImage(m, "hatified-", timestamp)
	saveResizedImage(m, "thumb-", timestamp, 0, 100)
}

func calculateHatRect(f Feature) image.Rectangle {
	d := distanceBetweenEyes(f)

	d2 := distanceBetween(f.LeftEye, f.Mouth)

	puts("distance between eyes:", d)
	puts("distance between left eye and mouth:", d2)

	x1 := f.LeftEye.X - d
	x2 := f.RightEye.X + d

	y1 := f.LeftEye.Y - d2 - ((x2 - x1) / 2)
	y2 := f.RightEye.Y + (d2 / 2)

	return image.Rect(x1, y1, x2, y2)
}

func distanceBetweenEyes(f Feature) int {
	return distanceBetween(f.LeftEye, f.RightEye)
}

func distanceBetween(p1, p2 image.Point) int {
	r := image.Rectangle{p1, p2}

	return int(math.Sqrt(
		math.Pow(float64(r.Size().X), 2) +
			math.Pow(float64(r.Size().Y), 2)))
}
