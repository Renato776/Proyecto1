package handles

import (
	sql2 "database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"test_servers/backend/db"
)

type Report struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	Carnet int `json:"carnet"`
	Course string `json:"course"`
	Body string `json:"body"`
	Server string `json:"server,omitempty"`
}

type ReportsCollection struct {
	Reports []Report `json:"reports"`
	Attended string `json:"attended"`
}

type IndividualReport struct {
	Report
	Attended string `json:"attended"`
}

type ErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func genericResponse (w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(data)
}

//Create a new report in DB
func CreateNewReport(w http.ResponseWriter, r *http.Request)  {
	decoder := json.NewDecoder(r.Body)
	var newReport Report
	err := decoder.Decode(&newReport)
	if err != nil {
		genericResponse(w, http.StatusBadRequest, ErrorResponse{http.StatusBadRequest, "Invalid json"})
		return
	}
	conn, _ := db.GetConnection()
	sql := "INSERT INTO Report(name, carnet, body, course, server) VALUES( $1, $2, $3, $4, $5 ) RETURNING id, carnet, name, body, course, server;"
	stmt, err := conn.Prepare(sql)
	if err != nil {
		fmt.Printf("Error on the insert: %v\n", err)
		genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Error on server"})
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(newReport.Name, newReport.Carnet, newReport.Body, newReport.Course, os.Getenv("ID")).Scan(&newReport.ID, &newReport.Carnet,
		&newReport.Name, &newReport.Body, &newReport.Course, &newReport.Server)
	if err != nil {
		fmt.Printf("Error on scan: %v\n", err)
		genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Error on server"})
		return
	}
	genericResponse(w, http.StatusCreated, newReport)
}

//GetAll or one report
func GetReports(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	conn, _ := db.GetConnection()
	if _, ok := vars["id"]; ok {
		sql := `SELECT id, carnet, name, body, course, server
				FROM Report
				WHERE id = $1;`
		stmt, err := conn.Prepare(sql)
		if err != nil {
			fmt.Printf("Error on the select: %v\n", err)
			genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Error on server"})
			return
		}
		defer stmt.Close()
		var report Report
		err = stmt.QueryRow(vars["id"]).Scan(&report.ID, &report.Carnet, &report.Name, &report.Body, &report.Course, &report.Server)
		if err != nil {
			fmt.Printf("Error on scan: %v\n", err)
			genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Error on server"})
			return
		}
		genericResponse(w, http.StatusOK, IndividualReport{report, fmt.Sprintf("Solicitud atendida por el servidor \"%s\"", os.Getenv("ID"))})
		return
	}
	sql := `SELECT id, carnet, name, body, course, server
			FROM Report;`
	stmt, err := conn.Prepare(sql)
	if err != nil {
		fmt.Printf("Error on the select: %v\n", err)
		genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Error on server"})
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		fmt.Printf("Error on the select: %v\n", err)
		genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Error on server"})
		return
	}
	defer rows.Close()
	var reports []Report = make([]Report, 0)
	for rows.Next() {
		var report Report
		if err := rows.Scan(&report.ID, &report.Carnet, &report.Name, &report.Body, &report.Course, &report.Server); err != nil {
			fmt.Printf("Error on scan: %v\n", err)
			genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Error on server"})
			return
		}
		reports = append(reports, report)
	}
	genericResponse(w, http.StatusOK, ReportsCollection{reports, fmt.Sprintf("Solicitud atendida por el servidor \"%s\"", os.Getenv("ID"))})
}

//Get Reports by Carnet
func GetReportsByCarnet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	conn, _ := db.GetConnection()
	var sql string
	var rows *sql2.Rows
	if carn, ok := vars["carnet"]; ok && carn != "" {
		sql = `SELECT id, carnet, name, body, course, server
				FROM Report
				WHERE carnet = $1`
		stmt, err := conn.Prepare(sql)
		if err != nil {
			fmt.Printf("Error on the select: %v\n", err)
			genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Error on server"})
			return
		}
		defer stmt.Close()
		rows, err = stmt.Query(vars["carnet"])
		if err != nil {
			fmt.Printf("Error on the select: %v\n", err)
			genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Error on server"})
			return
		}
	}else{
		sql := `SELECT id, carnet, name, body, course, server
				FROM Report;`
		stmt, err := conn.Prepare(sql)
		if err != nil {
			fmt.Printf("Error on the select: %v\n", err)
			genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Error on server"})
			return
		}
		defer stmt.Close()
		rows, err = stmt.Query()
		if err != nil {
			fmt.Printf("Error on the select: %v\n", err)
			genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Error on server"})
			return
		}
	}
	defer rows.Close()
	var reports []Report = make([]Report, 0)
	for rows.Next() {
		var report Report
		if err := rows.Scan(&report.ID, &report.Carnet, &report.Name, &report.Body, &report.Course, &report.Server); err != nil {
			fmt.Printf("Error on scan: %v\n", err)
			genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Error on server"})
			return
		}
		reports = append(reports, report)
	}
	genericResponse(w, http.StatusOK, ReportsCollection{reports, fmt.Sprintf("Solicitud atendida por el servidor \"%s\"", os.Getenv("ID"))})
}
