package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/thejohnny5/se_organization/pkg/models"
)

type FileRequestType struct {
	ID               uint   `json:"id"`
	OriginalFileName string `json:"original_file_name"`
	TypeOfDocument   string `json:"type_of_document"`
	DocumentName     string `json:"document_name"`
	Notes            string `json:"notes"`
	DownloadPath     string `json:"download_path"`
}

type FileDBHandler struct {
	DB *models.DBClient
}

func CreateFileDB(db *models.DBClient) *FileDBHandler {
	return &FileDBHandler{DB: db}
}

const MAX_UPLOAD_SIZE int64 = 1024 * 1024 * 250 // 250 Mb
const DIR_PATH string = "uploads"

func (db *FileDBHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
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
		log.Println("NO FILE")
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
	result := db.DB.DB.Create(&document)
	if result.Error != nil {
		http.Error(w, "Error writing to database", http.StatusInternalServerError)
		return
	}

	// NEED LOGIC TO HANDLE CASE WHERE STORING METADATA FAILS

	fmt.Fprintf(w, "Upload successful")
}

func (db *FileDBHandler) GetDocuments(w http.ResponseWriter, r *http.Request) {
	claims, err := GetClaims(r)

	if err != nil {
		http.Error(w, "Error with claims", http.StatusInternalServerError)
		return
	}
	// Query database for documents
	var docs []models.Document
	result := db.DB.DB.Where("user_id=?", claims.UserID).Find(&docs)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	docsReturn := make([]FileRequestType, len(docs))

	for i, doc := range docs {
		docsReturn[i] = FileRequestType{
			ID:               doc.ID,
			OriginalFileName: doc.OriginalFileName,
			TypeOfDocument:   doc.TypeOfDocument,
			Notes:            doc.Notes,
			DownloadPath:     fmt.Sprintf("/api/document/%v/download", doc.ID),
		}
	}

	// encode and return results
	json.NewEncoder(w).Encode(docsReturn)
}

// TODO: handle for only PDF
func (db *FileDBHandler) DownloadDoc(w http.ResponseWriter, r *http.Request) {
	claims, err := GetClaims(r)

	if err != nil {
		http.Error(w, "Error with claims", http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	docId := vars["id"] // query db to get url
	// cast docId to uint
	id, err := strconv.ParseUint(docId, 10, 32)
	if err != nil {
		http.Error(w, "Inproper ID", http.StatusBadRequest)
		return
	}
	doc := models.Document{ID: uint(id)}
	result := db.DB.DB.Where("user_id = ?", claims.UserID).First(&doc)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// path stored in Path -> send stream back
	f, err := os.Open(doc.Path)
	if err != nil {
		log.Printf("error opening: %s", doc.Path)
		http.Error(w, "Couldn't get file", http.StatusInternalServerError)
		return
	}
	defer f.Close()
	w.Header().Set("Content-type", "application/pdf")

	// reponse stream
	if _, err := io.Copy(w, f); err != nil {
		log.Printf("error streaming: %s", err)
		http.Error(w, "Couldn't stream file", http.StatusInternalServerError)
		return
	}
}

// func (db *DBClient) QueryTypeOfDoc(w http.ResponseWriter, r *http.Request) {
// 	// Read query string
// 	claims, err := GetClaims(r)

// 	if err != nil {
// 		http.Error(w, "Error with claims", http.StatusInternalServerError)
// 		return
// 	}

// 	query := r.URL.Query().Get("type_of_document")
// 	var files []models.Document
// 	result := db.DB.Where("user_id = ? && type_of_document = ?", claims.UserID, query).Find(files)

// 	if result.Error != nil {
// 		http.Error(w, result.Error.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(files)
// }
