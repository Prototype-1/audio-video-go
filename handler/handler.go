package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"encoding/json"
)

const (
	uploadDir = "assets"
	hlsDir    = "hls"
)

func RenderHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("media")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filePath := filepath.Join(uploadDir, header.Filename)
	outFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = outFile.ReadFrom(file)
	if err != nil {
		http.Error(w, "Failed to write file", http.StatusInternalServerError)
		return
	}

	go convertToHLS(filePath, header.Filename)

	fmt.Fprintf(w, "File uploaded successfully! Processing HLS conversion...\n")
}

func convertToHLS(inputPath, filename string) {
	outputFolder := filepath.Join(hlsDir, filename)
	os.MkdirAll(outputFolder, os.ModePerm)

	outputPath := filepath.Join(outputFolder, "index.m3u8")

	cmd := exec.Command("ffmpeg", "-i", inputPath, "-c:v", "libx264", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", outputPath)
	err := cmd.Run()
	if err != nil {
		log.Println("FFmpeg error:", err)
	} else {
		log.Println("HLS conversion completed for", filename)
	}
}

func GetLatestVideo(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(hlsDir)
	if err != nil {
		http.Error(w, "Failed to read HLS directory", http.StatusInternalServerError)
		return
	}

	var latestFolder string
	for _, file := range files {
		if file.IsDir() {
			latestFolder = file.Name()
		}
	}

	if latestFolder == "" {
		http.Error(w, "No processed video found", http.StatusNotFound)
		return
	}
	response := map[string]string{"videoPath": "/hls/" + latestFolder + "/index.m3u8"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

