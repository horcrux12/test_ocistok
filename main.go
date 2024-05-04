package main

import (
	"kalkulator_pinjaman/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/calc", controller.CalculatorController)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
