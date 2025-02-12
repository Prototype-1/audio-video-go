package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/Prototype-1/audio-video-go/handler"
)

const (
	uploadDir = "assets"
	hlsDir    = "hls"
)

func main() {
	os.MkdirAll(uploadDir, os.ModePerm)
	os.MkdirAll(hlsDir, os.ModePerm)

	http.HandleFunc("/", handler.RenderHomePage)
	http.HandleFunc("/upload", handler.HandleUpload)
	http.HandleFunc("/latest-video", handler.GetLatestVideo) 
	http.Handle("/hls/", http.StripPrefix("/hls/", http.FileServer(http.Dir(hlsDir))))

	fmt.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

