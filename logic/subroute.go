package logic

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
)

func AllTeams(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("template/index.html"))
	data := PageData{
		PageTitle: "List of Teams",
		Teams:     teams,
	}
	temp.Execute(w, data)
}

func GetTeam(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("template/team.html"))
	vars := mux.Vars(r)
	id := vars["id"]
	if _, ok := teamMap[id]; ok {
		data := PageData{
			PageTitle: teamMap[id].Name,
			Teams: []Team{
				teamMap[id],
			},
		}
		temp.Execute(w, data)
	} else {
		fmt.Fprintf(w, "Invalid team id: %s\n", id)
	}
}

func AddTeam(w http.ResponseWriter, r *http.Request) {
	var team Team
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &team); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	t := RepoCreateTeam(team)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
	json.NewDecoder(r.Body).Decode(&team)
	json.NewDecoder(r.Body).Decode(&team.Info)
	fmt.Fprintf(w, "Added %s (%s, %d and %s) to the Team repository.\n",
		team.Name, team.Info.Coach, team.Info.NoOfTitles, team.Info.Stadium)
}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {
}
