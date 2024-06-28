package server

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type TransferRecord struct {
	Filename  string
	Timestamp time.Time
	Success   bool
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	out, err := os.Create(filepath.Join("uploads", header.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	recordTransfer(header.Filename, err == nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("File uploaded successfully"))
}

func recordTransfer(filename string, success bool) {
	historyMutex.Lock()
	defer historyMutex.Unlock()

	transferHistory = append(transferHistory, TransferRecord{
		Filename:  filename,
		Timestamp: time.Now(),
		Success:   success,
	})
}
