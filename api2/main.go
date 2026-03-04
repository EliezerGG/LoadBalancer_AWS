package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Instancia  string `json:"Instancia"`
	Curso      string `json:"Curso"`
	Estudiante string `json:"Estudiante"`
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Instancia:  "Instancia #2 - API #2",
		Curso:      "Seminario de Sistemas 1",
		Estudiante: "Eliezer Guevara - 202100179",
	})
}

func main() {
	http.HandleFunc("/check", checkHandler)
	http.HandleFunc("/", rootHandler)

	fmt.Println("API #2 (Go) corriendo en puerto 3000")
	http.ListenAndServe(":3000", nil)
}