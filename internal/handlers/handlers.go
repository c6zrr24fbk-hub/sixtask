package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Metod not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusInternalServerError)
		return
	}

	myFile, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusInternalServerError)
		return
	}
	defer myFile.Close()

	data, err := io.ReadAll(myFile)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	content := string(data)
	result, err := service.AutoDetectAndConvert(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05")
	ext := filepath.Ext(header.Filename)
	if ext == "" {
		ext = ".txt"
	}
	outputFilename := "result_" + timestamp + ext
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		http.Error(w, "Unable to create result file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(result)
	if err != nil {
		http.Error(w, "Unable to write to result file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, result)
}
