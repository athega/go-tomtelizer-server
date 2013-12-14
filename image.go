package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	_ "image/png"
	"os"
)

func saveResizedImage(img image.Image, version, timestamp string, x, y int) {
	saveImage(resizeImage(img, x, y), version, timestamp)
}

func saveImage(img image.Image, version, timestamp string) {
	path := imageFileName(version, timestamp)

	file, _ := os.Create(path)

	defer file.Close()

	jpeg.Encode(file, img, &jpeg.Options{100})
}

func imageFileName(version, timestamp string) string {
	timestamp2 := "foo"
	return fmt.Sprintf("%s/%s-%s.jpg", UPLOAD_DIR, timestamp2, version)
}

func loadAndResizeImage(fn string, x, y int) image.Image {
	return resizeImage(loadImage(fn), x, y)
}

func resizeImage(img image.Image, x, y int) image.Image {
	return resize.Resize(uint(x), uint(y), img, resize.NearestNeighbor)
}

func loadImage(fn string) image.Image {
	file, err := os.Open(fn)
	defer file.Close()

	if err != nil {
		fatal(fmt.Sprintf("No such image: '%s'", fn))
	}
	img, _, err := image.Decode(file)

	if err != nil {
		fatal("Invalid image:", fn)
	} else {
		puts("Loaded image:", fn)
	}

	return img
}
