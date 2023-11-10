package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/thejohnny5/se_organization/pkg/models"
)

var headers []string = []string{"id", "company", "title", "location", "application_status",
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

func convertToJobStruct(dbRecord map[string]string) *models.JobApplication {
	jobApp := &models.JobApplication{}
	for key, value := range dbRecord {
		switch key {
		case "company":
			jobApp.Company = value
		case "title":
			jobApp.Title = value
		case "location":
			jobApp.Location = &value
		case "application_status":
			jobApp.ApplicationStatusId = nil
		case "application_type":
			jobApp.ApplicationTypeId = nil
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
			safeDeref(job.Location),
			job.ApplicationStatus.Text,
			job.ApplicationType.Text,
			job.Resume.DocumentName,
			job.CoverLetter.DocumentName,
			safeDeref(job.PostingUrl),
			safeDerefInt(job.SalaryLow),
			safeDerefInt(job.SalaryHigh),
			safeDeref(job.Notes),
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

func WriteJobsToDB(f *csv.Reader, mappings []JobHeaderMap, userId *uint) error {
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
	log.Printf("DBColumn: %+v", dbColumnIndices)
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

		job := convertToJobStruct(dbRecord)
		job.UserID = *userId
		log.Printf("DBRecord: %+v", job)
	}
	return nil
}
