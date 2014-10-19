package main

import (
	"net/http"
	"os"
)

func main() {
	setup()

	http.HandleFunc("/images/new", handleUpload)
	http.HandleFunc("/images/latest", handleLatestImages)
	http.HandleFunc("/images/latest.json", handleLatestImagesJSON)

	http.Handle("/uploaded_images/",
		http.StripPrefix("/uploaded_images/",
			http.FileServer(http.Dir(UploadDir))))

	port := getenv("PORT", "8080")

	fatal(http.ListenAndServe(":"+port, nil))
}

func getenv(key, fallback string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return fallback
}
