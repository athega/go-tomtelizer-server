package main

import (
	"image"
	"os"

	"github.com/joho/godotenv"
)

const (
	// MaxMemory defines a 1 MB Upload buffer
	MaxMemory = 1 * 1024 * 1024
)

var (
	// UploadDir is the upload directory
	UploadDir string

	// BaseURL contains the Tomtelizer base URL
	BaseURL string

	// SantaHat contains the image of the santa hat
	SantaHat image.Image

	// Debug flag
	Debug bool
)

func setup() {
	err := godotenv.Load()
	if err != nil {
		fatal("Error loading .env file")
	}

	// Debug mode
	Debug = os.Getenv("TOMTELIZER_DEBUG") == "true"
	puts("Debug:", Debug)

	UploadDir = os.Getenv("TOMTELIZER_UPLOAD_DIR")
	BaseURL = os.Getenv("TOMTELIZER_BASE_URL")
	SantaHat = loadImage(os.Getenv("TOMTELIZER_SANTA_HAT"))
}
