package logic

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", AllTeams).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", GetTeam).Methods("GET")
	// curl -H "Content-Type: application/json" -d '{"Name": "Dallas Cowboys", "Info": {"Coach": "Jason Garrett", "NoOfTitles": 5, "Stadium": "AT&T Stadium"}}' http://localhost:8080/add-team
	r.HandleFunc("/add-team", AddTeam).Methods("POST")
	fileServer()
	return r
}

func fileServer() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}
