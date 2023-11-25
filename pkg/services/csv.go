package services

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/thejohnny5/se_organization/pkg/models"
)

var headers []string = []string{"id", "date_applied", "company", "title", "location", "application_status",
	"application_type", "resume", "cover_letter", "posting_url", "salary_low", "salary_high", "notes"}
var set = genSet(headers)

func genSet(arr []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, header := range headers {
		set[header] = struct{}{}
	}
	return set
}

//	type JobHeaderMap struct {
//		Company           string `json:"company"`
//		Title             string `json:"title"`
//		Location          string `json:"location"`
//		ApplicationStatus string `json:"application_status"`
//		ApplicationType   string `json:"application_type"`
//		Resume            string `json:"resume"`
//		CoverLetter       string `json:"cover_letter"`
//		PostingUrl        string `json:"posting_url"`
//		SalaryRange       string `json:"salary_range"`
//		Notes             string `json:"notes"`
//	}
type JobHeaderMap struct {
	CSVHeader string `json:"csvHeader"`
	DBField   string `json:"dbField"`
}

func convertToJobStruct(dbRecord map[string]string, app_status_dd []models.Dropdown, app_type_dd []models.Dropdown) *models.JobApplication {
	jobApp := &models.JobApplication{}
	for key, value := range dbRecord {
		switch key {
		case "date_applied":
			// convert to timestamp
			t, err := time.Parse("10/20/2001", value)
			if err != nil {
				log.Printf("Error parsing date for val: %s", value)
				continue
			}
			jobApp.DateApplied = &t
		case "company":
			jobApp.Company = value
		case "title":
			jobApp.Title = value
		case "location":
			jobApp.Location = value
		case "application_status":
			id, err := parseDropdown(value, app_status_dd)
			log.Printf("Value: %s, id: %d", value, id)
			if err != nil {
				log.Printf("Error parsing uint for val: %s", value)
				continue
			}
			jobApp.ApplicationStatusId = &id
		case "application_type":
			id, err := parseDropdown(value, app_type_dd)
			if err != nil {
				log.Printf("Error parsing uint for val: %s", value)
				continue
			}
			jobApp.ApplicationTypeId = &id
		case "posting_url":
			jobApp.PostingUrl = value
		case "notes":
			jobApp.Notes = value
		default:
			log.Printf("None for key: %s", key)
		}
	}
	return jobApp
}

func WriteJobsToClient(w http.ResponseWriter, jobs *[]models.JobApplication) {
	csvWriter := csv.NewWriter(w)
	csvWriter.Write(headers)
	for _, job := range *jobs {
		line := []string{
			strconv.Itoa(int(job.ID)),
			job.Company,
			job.Title,
			job.Location,
			job.ApplicationStatus.Text,
			job.ApplicationType.Text,
			job.Resume.DocumentName,
			job.CoverLetter.DocumentName,
			job.PostingUrl,
			strconv.Itoa(job.SalaryLow),
			strconv.Itoa(job.SalaryHigh),
			job.Notes,
		}
		if err := csvWriter.Write(line); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	// Flush buffered data
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func safeDerefInt(i *int) string {
	if i != nil {
		return strconv.Itoa(*i)
	}
	return ""
}

func safeDeref(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func WriteJobsToDB(f *csv.Reader, mappings []JobHeaderMap, userId *uint, db *models.DBClient) error {
	// Map our first csv row headers to the appropriate index in array
	csvHeaders, err := f.Read()
	if err != nil {
		fmt.Println("Error opening")
		return err
	}
	headerIndices := make(map[string]int)
	for i, header := range csvHeaders {
		headerIndices[header] = i
	}

	// Translate this to our map //also ensure that it is in our set
	dbColumnIndices := make(map[string]int)
	for _, mapping := range mappings {
		if _, ok := set[mapping.DBField]; ok {
			if index, ok := headerIndices[mapping.CSVHeader]; ok {
				dbColumnIndices[mapping.DBField] = index
			} else {
				log.Printf("csv header %s not found", mapping.CSVHeader)
			}
		}
	}

	// Get dropdown options for application status and type
	var app_stat_dd []models.Dropdown
	var app_type_dd []models.Dropdown
	db.DB.Where("user_id is NULL").Where("table_type = ?", "application_status").Find(&app_stat_dd)
	db.DB.Where("user_id is NULL").Where("table_type = ?", "application_type").Find(&app_type_dd)

	// Read through rest of rows and apply the mappings
	for {
		record, err := f.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error reading from csv: %v", err)
		}

		dbRecord := make(map[string]string)
		for dbField, index := range dbColumnIndices {
			dbRecord[dbField] = record[index]
		}

		job := convertToJobStruct(dbRecord, app_stat_dd, app_type_dd)
		job.UserID = *userId
		db.DB.Save(&job)
	}
	return nil
}

func parseDropdown(s string, dd []models.Dropdown) (uint, error) {
	// Try to match the application status to number in DB
	for _, drop := range dd {
		// try to match drop text
		if strings.EqualFold(s, drop.Text) {
			return drop.ID, nil
		}
	}
	// if No match, returns 0
	return 0, errors.New("improper value for s")
}
