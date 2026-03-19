package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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


func getDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data := map[string]string{
		"Instancia":  "Instancia #2 - API #2",
		"Curso":      "Seminario de Sistemas 1",
		"Estudiante": "Eliezer Guevara - 202100179",
		"Lenguaje":   "Go",
		"Puerto":     getPort(),
	}
	json.NewEncoder(w).Encode(data)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"status":"ok","api":"API-2"}`)
}

func getPort() string {
	if p := os.Getenv("PORT"); p != "" {
		return p
	}
	return "3000"
}

func main() {

    http.HandleFunc("/get-data", getDataHandler)
    http.HandleFunc("/health",   healthHandler)
	http.HandleFunc("/check", checkHandler)
	http.HandleFunc("/", rootHandler)

	fmt.Println("API #2 (Go) corriendo en puerto 3000")
	http.ListenAndServe(":3000", nil)
}