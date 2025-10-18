package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CatFactStruct struct {
	Fact string `json:"fact"`
}

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Stack string `json:"stack"`
}

type ApiResponse struct {
	Status string `json:"status"`
	User User `json:"user"`
	Timestamp string `json:"timestamp"`
	Fact CatFactStruct `json:"fact"`
}

func meHandler(w http.ResponseWriter, r *http.Request) {
	resp, err  := http.Get("https://catfact.ninja/fact")

	if err != nil {
		http.Error(w, "failed to fetch from cat fact", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()


	var catFact CatFactStruct
	if err := json.NewDecoder(resp.Body).Decode(&catFact); err != nil {
		http.Error(w,"Failed to parse cat fact", http.StatusInternalServerError)
		return
	}
	

	response := ApiResponse {
		Status: "success",
		User: User{
			Name: "Buhari Ahmed Alhassan",
			Email: "buhari.alhassan0@gmail.com",
			Stack: "Backend Golang",
		},
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Fact: catFact,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/me", meHandler)
	fmt.Println("Server running")
	http.ListenAndServe(":8080",nil)

}