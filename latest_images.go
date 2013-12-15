package main

import (
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

	fmt.Fprintln(w, latestImagesXML())
}

func marshalXML(images []*Image) string {
	i := Images{Type: "array", Image: images}

	out, err := xml.MarshalIndent(i, "", "  ")

	if err != nil {
		puts("error: %v\n", err)
	}

	return xml.Header + string(out)
}

func latestImagesXML() string {
	images := handleFileStats(func(f os.FileInfo, id int) *Image {
		return &Image{
			Id:        atoi(strings.TrimRight(f.Name(), ".jpg")),
			Filename:  f.Name(),
			CreatedAt: f.ModTime(),
			Size:      f.Size(),
			Checksum:  "foo"}
	})

	return marshalXML(images)
}

type Image struct {
	XMLName     xml.Name  `xml:"image"`
	Id          int       `xml:"id"`
	Filename    string    `xml:"filename"`
	Orientation int       `xml:"orientation"`
	CreatedAt   time.Time `xml:"created-at"`
	UpdatedAt   time.Time `xml:"updated-at"`
	Width       int       `xml:"width"`
	Height      int       `xml:"height"`
	Size        int64     `xml:"hatified-file-size"`
	Checksum    string    `xml:"hatified-file-checksum"`
}

type Images struct {
	XMLName xml.Name `xml:"images"`
	Type    string   `xml:"type,attr"`
	Image   []*Image
}

type fileStatsHandler func(os.FileInfo, int) *Image

func handleFileStats(yield fileStatsHandler) []*Image {
	allFileNames, _ := filepath.Glob(UPLOAD_DIR + "/*.jpg")

	fileNames := []string{}

	for _, fn := range allFileNames {
		if strings.Contains(fn, "-") == false {
			fileNames = append(fileNames, fn)
		}
	}

	i := []*Image{}

	for idx, fn := range fileNames {
		f, err := os.Stat(fn)

		if err != nil {
			puts("Unable to stat", fn)
		}

		i = append(i, yield(f, idx+1))
	}

	return i
}

func contentType(w http.ResponseWriter, t string) {
	w.Header().Set("Content-Type", t)
}
