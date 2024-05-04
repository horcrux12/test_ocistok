package controller

import (
	"encoding/json"
	"kalkulator_pinjaman/lib/constanta"
	"kalkulator_pinjaman/model"
	"kalkulator_pinjaman/service"
	"log"
	"net/http"
	"strings"
	"time"
)

func CalculatorController(rw http.ResponseWriter, r *http.Request) {
	var (
		payload  model.CalculatorLoanRequest
		response []model.LoanTable
	)

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		writeResponseBody(rw, map[string]interface{}{
			"message": "Internal Server Error",
		}, http.StatusInternalServerError)
		return
	}

	payload.StartDate, err = time.Parse(constanta.YYYYMMDD, strings.ReplaceAll(payload.StartDateStr, " ", ""))
	if err != nil {
		log.Println(err)
		writeResponseBody(rw, map[string]interface{}{
			"message": "Internal Server Error",
		}, http.StatusInternalServerError)
		return
	}

	response, err = service.CalculateLoan(r.Context(), payload)
	if err != nil {
		log.Println(err)
		writeResponseBody(rw, map[string]interface{}{
			"message": "Internal Server Error",
		}, http.StatusInternalServerError)
		return
	}

	writeResponseBody(rw, map[string]interface{}{
		"message": "Success",
		"data":    response,
	}, http.StatusOK)
}

func writeResponseBody(rw http.ResponseWriter, data interface{}, status int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(data)
}
