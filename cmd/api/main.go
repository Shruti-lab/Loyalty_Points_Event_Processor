// package api
package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=user password=pass dbname=loyalty sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/points/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID := vars["user_id"]

		var points int
		err := db.QueryRow(`SELECT points FROM user_points WHERE user_id = $1`, userID).Scan(&points)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"user_id": userID,
			"points":  points,
		})
	}).Methods("GET")

	log.Println("API running on :8080")
	http.ListenAndServe(":8080", r)
}
