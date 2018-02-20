package main

import (
	"log"
	"net/http"

	// Note this is my path according to my GOPATH, chage it according to yours.
	"./controllers"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/upload", controllers.UploadFile)
	log.Println("Running")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
