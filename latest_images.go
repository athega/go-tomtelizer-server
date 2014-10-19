package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func handleLatestImages(w http.ResponseWriter, r *http.Request) {
	contentType(w, "application/xml; charset=utf-8")

	fmt.Fprintln(w, marshalXML(latestImages()))
}

func handleLatestImagesJSON(w http.ResponseWriter, r *http.Request) {
	contentType(w, "application/json; charset=utf-8")

	fmt.Fprintln(w, marshalJSON(latestImages()))
}

func imageURL(fn string) string {
	return BaseURL + "/uploaded_images/" + fn
}

func latestImages() Images {
	images := handleFileStats(func(f os.FileInfo) *Image {
		idStr := strings.TrimRight(f.Name(), ".jpg")
		return &Image{
			ID:        atoi(idStr),
			Filename:  f.Name(),
			CreatedAt: f.ModTime(),
			UpdatedAt: f.ModTime(),
			Size:      f.Size(),
			URL:       imageURL(f.Name())}
	})

	return Images{Type: "array", Image: images}
}

// Image represents the data for an image
type Image struct {
	XMLName   xml.Name  `xml:"image" json:"-"`
	ID        int       `xml:"id" json:"id"`
	Filename  string    `xml:"filename" json:"filename"`
	CreatedAt time.Time `xml:"created-at" json:"created_at"`
	UpdatedAt time.Time `xml:"updated-at" json:"updated_at"`
	Size      int64     `xml:"hatified-file-size"`
	URL       string    `xml:"url"`
}

// Images represent a list of images
type Images struct {
	XMLName xml.Name `xml:"images" json:"-"`
	Type    string   `xml:"type,attr" json:"-"`
	Image   []*Image `json:"images"`
}

type fileStatsHandler func(os.FileInfo) *Image

func handleFileStats(yield fileStatsHandler) []*Image {
	allFileNames, _ := filepath.Glob(UploadDir + "/*.jpg")

	fileNames := []string{}

	for _, fn := range allFileNames {
		if strings.Contains(fn, "-") == false {
			fileNames = append(fileNames, fn)
		}
	}

	i := []*Image{}

	for _, fn := range fileNames {
		f, err := os.Stat(fn)

		if err != nil {
			puts("Unable to stat", fn)
		}

		i = append(i, yield(f))
	}

	return i
}

func contentType(w http.ResponseWriter, t string) {
	w.Header().Set("Content-Type", t)
}

func marshalXML(images Images) string {
	out, err := xml.MarshalIndent(images, "", "  ")

	if err != nil {
		puts("error: %v\n", err)
	}

	return xml.Header + string(out)
}

func marshalJSON(images Images) string {
	out, err := json.MarshalIndent(images, "", "  ")

	if err != nil {
		puts("error: %v\n", err)
	}

	return string(out)
}
