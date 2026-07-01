package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sanrose-bhets/Job-Grinder/internal/database"
)

// a http handler that the go standard library expects
func (apiCfg *apiConfig) interviewCreateHandler(w http.ResponseWriter, r *http.Request) {

	type parameter struct {
		Status        string    `json:"status"`
		InterviewDate time.Time `json:"interviewDate"`
	}

	decode := json.NewDecoder(r.Body)

	param := parameter{}
	err := decode.Decode(&param)
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error parsing json body: %v", err))
		return
	}

	comp, err := apiCfg.DB.CreateInterviews(r.Context(), database.CreateInterviewsParams{
		Status:        param.Status,
		Interviewdate: param.InterviewDate,
	})
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Unable to create interview: %v", err))
		return
	}

	respondwithJson(w, 200, databaseIntvwtoIntvw(comp))

}
