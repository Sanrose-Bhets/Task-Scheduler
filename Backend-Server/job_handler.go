package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sanrose-bhets/Job-Grinder/internal/database"
)

// a http handler that the go standard library expects
func (apiCfg *apiConfig) jobCreateHandler(w http.ResponseWriter, r *http.Request) {

	type parameter struct {
		Positionname         string `json:"name"`
		Companyid            int32  `json:"companyid"`
		Experiencerequired   string `json:"experienceRequired"`
		Expectedcompensation string `json:"expectedCompensation"`
		Employmenttype       string `json:"employmentType"`
		Joblocationtype      string `json:"jobLocationType"`
	}

	decode := json.NewDecoder(r.Body)

	param := parameter{}
	err := decode.Decode(&param)
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error parsing json body: %v", err))
		return
	}

	job, err := apiCfg.DB.CreateJobs(r.Context(), database.CreateJobsParams{
		Positionname:         param.Positionname,
		Companyid:            param.Companyid,
		Experiencerequired:   param.Experiencerequired,
		Expectedcompensation: param.Expectedcompensation,
		Employmenttype:       param.Employmenttype,
		Joblocationtype:      param.Joblocationtype,
	})
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Unable to create job: %v", err))
		return
	}

	respondwithJson(w, 200, databaseJobtoJob(job))

}
