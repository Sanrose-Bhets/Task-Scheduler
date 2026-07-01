package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sanrose-bhets/Job-Grinder/internal/database"
)

// a http handler that the go standard library expects
func (apiCfg *apiConfig) applicationCreateHandler(w http.ResponseWriter, r *http.Request) {

	type parameter struct {
		Department  string `json:"department"`
		Status      string `json:"status"`
		Company     int32  `json:"company"`
		Position    int32  `json:"position"`
		InterviewId int32  `json:"interviewId"`
	}

	decode := json.NewDecoder(r.Body)

	param := parameter{}
	err := decode.Decode(&param)
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error parsing json body: %v", err))
		return
	}

	app, err := apiCfg.DB.CreateApplications(r.Context(), database.CreateApplicationsParams{
		Department: param.Department,
		Status:     param.Status,
		Company:    param.Company,
		Position:   param.Position,
		Interviewid: sql.NullInt32{
			Int32: param.InterviewId,
			Valid: param.InterviewId != 0,
		},
	})
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Unable to create application: %v", err))
		return
	}

	respondwithJson(w, 200, databaseAppToApp(app))

}
