package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ryanyogan/goblog/accountservice/dbclient"
)

// DBClient defines a single instance of the Bolt DB Instance
var DBClient dbclient.IBoltClient

// GetAccount -
func GetAccount(w http.ResponseWriter, r *http.Request) {
	var accountID = mux.Vars(r)["accountID"]

	account, err := DBClient.QueryAccount(accountID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// HealthCheck --
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	dbUp := DBClient.Check()
	if dbUp {
		data, _ := json.Marshal(healthCheckResponse{Status: "UP"})
		writeJSONResponse(w, http.StatusOK, data)
	} else {
		data, _ := json.Marshal(healthCheckResponse{Status: "Database Down"})
		writeJSONResponse(w, http.StatusServiceUnavailable, data)
	}
}

func writeJSONResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}

type healthCheckResponse struct {
	Status string `json:"status"`
}
