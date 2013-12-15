package main

import "net/http"

func main() {
	setup()

	http.HandleFunc("/images/new", handleUpload)
	http.HandleFunc("/images/latest", handleLatestImages)

	http.Handle("/uploaded_images/",
		http.StripPrefix("/uploaded_images/",
			http.FileServer(http.Dir(UPLOAD_DIR))))

	fatal(http.ListenAndServe(":8080", nil))
}
