package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/thejohnny5/se_organization/pkg/models"
)

const MAX_UPLOAD_SIZE int64 = 1024 * 1024 * 250 // 250 Mb
const DIR_PATH string = "../../uploads"

func (db *DBClient) UploadFile(w http.ResponseWriter, r *http.Request) {
	// Get user id -> fail early
	claims, err := GetClaims(r)

	if err != nil {
		http.Error(w, "Error with claims", http.StatusInternalServerError)
		return
	}

	// Limits file size
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big.", http.StatusBadRequest)
		return
	}

	// Get releavant file info
	file_name := r.FormValue("original_file_name")
	document_name := r.FormValue("document_name")
	type_of_document := r.FormValue("type_of_document")
	notes := r.FormValue("notes")
	// Check for correct metadata

	// Get file info
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create folder if DNE
	err = os.MkdirAll(DIR_PATH, os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filePath := fmt.Sprintf("%s/%d%s", DIR_PATH, time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))

	// Create a new file in directory
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy file contents to location
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// IF successful, we can write the file to the database
	// Construct model
	document := models.Document{UserID: claims.UserID, OriginalFileName: file_name, Path: filePath, TypeOfDocument: type_of_document, DocumentName: document_name, Notes: notes}
	result := db.DB.Create(&document)
	if result.Error != nil {
		http.Error(w, "Error writing to database", http.StatusInternalServerError)
		return
	}

	// NEED LOGIC TO HANDLE CASE WHERE STORING METADATA FAILS

	fmt.Fprintf(w, "Upload successful")
}
