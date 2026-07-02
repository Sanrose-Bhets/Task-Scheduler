package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	authentication "github.com/sanrose-bhets/Job-Grinder/auth"
	"github.com/sanrose-bhets/Job-Grinder/internal/database"
)

// a http handler that the go standard library expects
func (apiCfg *apiConfig) accountCreateHandler(w http.ResponseWriter, r *http.Request) {

	type parameter struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Role  string `json:"role"`
	}

	decode := json.NewDecoder(r.Body)

	param := parameter{}
	err := decode.Decode(&param)
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error parsing json body: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name:  param.Name,
		Email: param.Email,
		Role:  param.Role,
	})
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Unable to create user: %v", err))
		return
	}

	respondwithJson(w, 200, databaseAcctoAcc(user))

}

func (apiCfg *apiConfig) accountGetByAPIHandler(w http.ResponseWriter, r *http.Request) {

	ApiKey, err := authentication.GetAPIKey(r.Header)
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Unable to get API key: %v", err))
		return
	}
	user, err := apiCfg.DB.GetUserbyAPI(r.Context(), ApiKey)
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Unable to get user: %v", err))
		return
	}
	respondwithJson(w, 200, databaseAcctoAcc(user))
}

func (apiCfg *apiConfig) accountGetAllHandler(w http.ResponseWriter, r *http.Request) {
	users, err := apiCfg.DB.GetAllUsers(r.Context())
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Unable to get users: %v", err))
		return
	}
	respondwithJson(w, 200, users)
}
