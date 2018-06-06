package logic

import (
	"fmt"
	"strconv"
)

var currentIndex int
var teamMap map[string]Team
var teams []Team

func init() {
	teamMap = make(map[string]Team)
	RepoCreateTeam(Team{Name: "New Orleans Saints", Info: TeamData{Coach: "Sean Payton", NoOfTitles: 1, Stadium: "SuperDome"}})
	RepoCreateTeam(Team{Name: "Cleveland Browns", Info: TeamData{Coach: "Hugh Jackson", NoOfTitles: 0, Stadium: "Dawg Pound"}})
	RepoCreateTeam(Team{Name: "New England Patriots", Info: TeamData{Coach: "Bill Belichick", NoOfTitles: 5, Stadium: "Gillette Stadium"}})
}

func RepoFindTeam(name string) Team {
	for _, l := range teamMap {
		if l.Name == name {
			return l
		}
	}
	return Team{}
}

func RepoCreateTeam(t Team) Team {
	currentIndex += 1
	t.Id = strconv.Itoa(currentIndex)
	teamMap[strconv.Itoa(currentIndex)] = t
	teams = append(teams, t)
	return t
}

func RepoDestroyTeam(id string) error {
	_, ok := teamMap[id]
	if ok {
		delete(teamMap, id)
		return nil
	}
	return fmt.Errorf("could not find team with id of %d to delete", id)
}
