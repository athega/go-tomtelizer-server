package main

import "net/http"

func main() {
	setup()

	http.HandleFunc("/images/new", handleUpload)
	http.Handle("/", http.FileServer(http.Dir(UPLOAD_DIR)))

	fatal(http.ListenAndServe(":8080", nil))
}
