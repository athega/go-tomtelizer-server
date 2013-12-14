package main

import (
	"github.com/joho/godotenv"
	"image"
	"os"
	"runtime"
)

const (
	// 1MB Upload buffer
	MAX_MEMORY = 1 * 1024 * 1024
)

var (
	UPLOAD_DIR string
	SANTA_HAT  image.Image
	DEBUG      bool
)

func setup() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	puts("Available CPUs:", runtime.NumCPU())

	err := godotenv.Load()
	if err != nil {
		fatal("Error loading .env file")
	}

	// Debug mode
	DEBUG = os.Getenv("TOMTELIZER_DEBUG") == "true"
	puts("DEBUG:", DEBUG)

	UPLOAD_DIR = os.Getenv("TOMTELIZER_UPLOAD_DIR")
	SANTA_HAT = loadImage(os.Getenv("TOMTELIZER_SANTA_HAT"))
}
