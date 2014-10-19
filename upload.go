package main

import (
	"image"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handleUpload(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	if err := r.ParseMultipartForm(MaxMemory); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	// Get a timestamp string used in the filename
	timestamp := timestampString()
	path := imageFileName("", timestamp)

	writeUploadedFile(r, path)
	processFeatures(r, path, timestamp)
}

func processFeatures(r *http.Request, path, timestamp string) {
	var preparedImage image.Image

	original := loadImage(path)

	// Check if we should process the features
	if r.FormValue("processFeatures") == "true" {
		puts("Processing", path)

		// Get the facial feature count
		ffc := len(r.Form["features[][mouth_x]"])

		features := make([]Feature, 0, ffc)

		h := original.Bounds().Max.Y

		for i := 0; i < ffc; i++ {
			lEyeX := atoi(r.Form["features[][left_eye_x]"][i])
			lEyeY := atoi(r.Form["features[][left_eye_y]"][i])
			rEyeX := atoi(r.Form["features[][right_eye_x]"][i])
			rEyeY := atoi(r.Form["features[][right_eye_y]"][i])
			mouthX := atoi(r.Form["features[][mouth_x]"][i])
			mouthY := atoi(r.Form["features[][mouth_y]"][i])

			features = append(features, Feature{
				image.Point{lEyeX, h - lEyeY},
				image.Point{rEyeX, h - rEyeY},
				image.Point{mouthX, h - mouthY},
			})
		}

		preparedImage = handleFacialFeatures(features, original, timestamp)
	} else {
		preparedImage = original
	}

	saveImage(preparedImage, "hatified-", timestamp)
	saveResizedImage(preparedImage, "thumb-", timestamp, 0, 200)
}

func writeUploadedFile(r *http.Request, imagePath string) {
	file, _, _ := r.FormFile("uploaded")
	defer file.Close()

	buf, _ := ioutil.ReadAll(file)
	ioutil.WriteFile(imagePath, buf, os.ModePerm)
}
