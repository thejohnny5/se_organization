package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/thejohnny5/se_organization/pkg/models"
	"github.com/thejohnny5/se_organization/pkg/services"
	"gorm.io/gorm"
)

// JobApplicationRequest represents the data structure for a job application request.
type JobApplicationRequest struct {
	ID                  uint   `json:"id"`
	DateApplied         string `json:"date_applied"`
	Company             string `json:"company"`
	Title               string `json:"title"`
	Location            string `json:"location"` // Pointer to support omission if empty
	ApplicationStatusId *uint  `json:"application_status_id"`
	ApplicationStatus   string `json:"application_status"` // FK to dropdowns
	ApplicationTypeId   *uint  `json:"application_type_id"`
	ApplicationType     string `json:"application_type"` // FK to dropdown
	ResumeId            *uint  `json:"resume_id"`        // Pointer to support omission if empty
	ResumeName          string `json:"resume_name"`
	ResumeDownload      string `json:"resume_download"`
	CoverLetterId       *uint  `json:"cover_letter_id"` // Pointer to support omission if empty
	CoverLetterName     string `json:"cover_letter_name"`
	CoverLetterDownload string `json:"cover_letter_download"`
	PostingUrl          string `json:"posting_url"`
	SalaryLow           int    `json:"salary_low"`
	SalaryHigh          int    `json:"salary_high"`
	Notes               string `json:"notes"`
}

// JobsDBHandler handles database operations related to job applications.
// It provides methods for creating, retrieving, updating, and deleting job applications.
type JobsDBHandler struct {
	DB *models.DBClient // DBClient instance for database operations
}

// CreateJobsDB initializes a new JobsDBHandler with the given DBClient.
// This function is typically called at the start of the application to set up
// a handler for job-related database operations.
func CreateJobsDB(db *models.DBClient) *JobsDBHandler {
	return &JobsDBHandler{DB: db}
}

// GetJobs retrieves a list of job applications associated with the user from the database.
// It extracts the user's claims from the request, queries the database for jobs tied to the user,
// and returns them in a JSON-encoded format. In case of an error, it logs the error and sends
// an appropriate HTTP response.
func (db *JobsDBHandler) GetJobs(w http.ResponseWriter, r *http.Request) {
	// Grab claims value from context
	claims, err := GetClaims(r)
	if err != nil {
		log.Printf("error returning claims: %s", err)
		http.Error(w, "Error with claims", http.StatusInternalServerError)
		return
	}
	// Get jobs associated with user id i in
	var jobs []models.JobApplication
	if err := db.PreloadDocument(claims, &jobs).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert jobs from models.JobApplication to JobApplicationRequest for JSON encoding
	jobs_json := make([]JobApplicationRequest, len(jobs))
	for i, job := range jobs {
		jobs_json[i] = JobApplicationRequest{
			ID:                  job.ID,
			DateApplied:         dateDeref(job.DateApplied),
			Company:             job.Company,
			Title:               job.Title,
			Location:            job.Location,
			ApplicationStatusId: job.ApplicationStatusId,
			ApplicationStatus:   job.ApplicationStatus.Text,
			ApplicationTypeId:   job.ApplicationTypeId,
			ApplicationType:     job.ApplicationType.Text,
			ResumeId:            job.ResumeId,
			ResumeName:          job.Resume.DocumentName,
			ResumeDownload:      fmt.Sprintf("/api/document/%d/download", job.ResumeId),
			CoverLetterId:       job.CoverLetterId,
			CoverLetterName:     job.CoverLetter.DocumentName,
			CoverLetterDownload: fmt.Sprintf("/api/document/%d/download", job.CoverLetterId),
			PostingUrl:          job.PostingUrl,
			SalaryLow:           job.SalaryLow,
			SalaryHigh:          job.SalaryHigh,
			Notes:               job.Notes,
		}
	}

	// Respond with the list of job applications in JSON format
	json.NewEncoder(w).Encode(jobs_json)
}

func (db *JobsDBHandler) CreateJob(w http.ResponseWriter, r *http.Request) {
	claims, err := GetClaims(r)
	if err != nil {
		http.Error(w, "Error with claims", http.StatusInternalServerError)
		return
	}
	// get request body decodeded
	body, berr := io.ReadAll(r.Body)
	defer r.Body.Close()
	if berr != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var job_json JobApplicationRequest

	if err := json.Unmarshal(body, &job_json); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// set user_id to claims.UserID (or could check if they match before deciding)
	job := models.JobApplication{
		UserID:              claims.UserID,
		Company:             job_json.Company,
		Title:               job_json.Title,
		Location:            job_json.Location,
		ApplicationStatusId: job_json.ApplicationStatusId,
		ApplicationTypeId:   job_json.ApplicationTypeId,
		ResumeId:            job_json.ResumeId,
		CoverLetterId:       job_json.CoverLetterId,
		PostingUrl:          job_json.PostingUrl,
		SalaryLow:           job_json.SalaryLow,
		SalaryHigh:          job_json.SalaryHigh,
		Notes:               job_json.Notes,
	}

	result := db.DB.DB.Create(&job)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	job_json.ID = job.ID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(job_json)
}

func (db *JobsDBHandler) UpdateJob(w http.ResponseWriter, r *http.Request) {
	claims, err := GetClaims(r)
	if err != nil {
		http.Error(w, "Error with claims", http.StatusInternalServerError)
		return
	}
	// get request body decodeded
	body, berr := io.ReadAll(r.Body)
	defer r.Body.Close()
	if berr != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var job_json JobApplicationRequest
	if err := json.Unmarshal(body, &job_json); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	job := models.JobApplication{
		UserID:              claims.UserID,
		DateApplied:         convertDate(job_json.DateApplied),
		ID:                  job_json.ID,
		Company:             job_json.Company,
		Title:               job_json.Title,
		Location:            job_json.Location,
		ApplicationStatusId: job_json.ApplicationStatusId,
		ApplicationTypeId:   job_json.ApplicationTypeId,
		ResumeId:            job_json.ResumeId,
		CoverLetterId:       job_json.CoverLetterId,
		PostingUrl:          job_json.PostingUrl,
		SalaryLow:           job_json.SalaryLow,
		SalaryHigh:          job_json.SalaryHigh,
		Notes:               job_json.Notes,
	}

	// set user_id to claims.UserID (or could check if they match before deciding)
	job.UserID = claims.UserID
	result := db.DB.DB.Save(&job)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (db *JobsDBHandler) DeleteJob(w http.ResponseWriter, r *http.Request) {
	claims, err := GetClaims(r)
	if err != nil {
		http.Error(w, "Error with claims", http.StatusInternalServerError)
		return
	}

	// extract vars
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is missing in parameters", http.StatusBadRequest)
		return
	}

	// Parse the ID to a uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// set user_id to claims.UserID (or could check if they match before deciding)
	recordToDelete := models.JobApplication{ID: uint(id), UserID: claims.UserID}

	result := db.DB.DB.Delete(&recordToDelete)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (db *JobsDBHandler) PreloadDocument(claims Claims, jobs *[]models.JobApplication) *gorm.DB {
	return db.DB.DB.Preload("ApplicationStatus").Preload("ApplicationType").
		Preload("Resume").Preload("CoverLetter").
		Where("user_id = ?", claims.UserID).
		Order("created_at DESC").Find(jobs) // Notice 'jobs' instead of '&jobs'
}

func (db *JobsDBHandler) DownloadCSV(w http.ResponseWriter, r *http.Request) {
	claims, err := GetClaims(r)
	if err != nil {
		http.Error(w, "Error with claims", http.StatusInternalServerError)
		return
	}
	// Get all Record for the current user
	var jobs []models.JobApplication
	if err := db.DB.DB.Where("user_id = ?", claims.UserID).Find(&jobs).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("jobs: +%v", jobs)
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=job-apps.csv")

	services.WriteJobsToClient(w, &jobs)
}

func (db *JobsDBHandler) UploadCSV(w http.ResponseWriter, r *http.Request) {
	claims, err := GetClaims(r)
	if err != nil {
		http.Error(w, "Error with claims", http.StatusInternalServerError)
		return
	}

	// Limit file size and parse form
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big.", http.StatusBadRequest)
		return
	}

	mappings := []services.JobHeaderMap{}
	mappingJSON := r.FormValue("headerMapping")

	log.Printf("MappingJSON: %v", mappingJSON)

	if err := json.Unmarshal([]byte(mappingJSON), &mappings); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("mappings: %+v", mappings)
	file, _, err := r.FormFile("uploadFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a CSV reader and read the first row (headers)
	csvReader := csv.NewReader(file)

	if err := services.WriteJobsToDB(csvReader, mappings, &claims.UserID, db.DB); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("CSV Uploaded"))

}

func dateDeref(d *time.Time) string {
	// If nil pointer return empty string
	if d == nil {
		return ""
	}
	return d.Format("Mon Jan 02 2006")
}

func convertDate(isoDateString string) *time.Time {
	t, err := time.Parse(time.RFC3339, isoDateString)
	if err != nil {
		log.Printf("Couldn't convert date string: %s", isoDateString)
		return nil
	}
	return &t
}
