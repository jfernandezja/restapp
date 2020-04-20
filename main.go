package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
)

type info struct {
	Version     string `json:"version"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Description string `json:"Description"`
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	var myInfo = info{
		Version:     "v2",
		Name:        "Juan Carlos Fernández Jara",
		Username:    "jfernandezja",
		Description: "UOC - PEC3 - GOLang Rest Application",
	}
	result := 0.0
	for i := math.Pow(8, 7); i >= 0; i-- {
		result += math.Atan(i) * math.Tan(i)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myInfo)
}

func main() {
	http.HandleFunc("/restapp/v2/info", getInfo)
	fmt.Println("Starting Rest APP on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
