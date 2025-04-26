package handlers

import (
	"fmt"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	"net/http"
	"os"
	"time"
)

func ReturnHTML(w http.ResponseWriter, r *http.Request) {
	filePath := "index.html"
	file, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}

func ParseHTML(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(10)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error getting file from form", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data := make([]byte, handler.Size)
	_, err = file.Read(data)
	if err != nil {
		http.Error(w, "Error getting file from form", http.StatusInternalServerError)
		return
	}

	result := service.CheckMorse(string(data))

	resultFile, err := os.OpenFile(time.Now().UTC().String(), os.O_CREATE|os.O_WRONLY, 0755)
	defer file.Close()
	_, err = fmt.Fprint(resultFile, result)
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
