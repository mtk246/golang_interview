package assignment3

import (
	"fmt"
	"sort"
)

type Team struct {
	Name        string
	Win         int
	Lose        int
	Draw        int
	GoalFor     int
	GoalAgainst int
}

func (t Team) Score() int {
	return t.Win*3 + t.Draw
}

func (t Team) GoalDiff() int {
	return t.GoalFor - t.GoalAgainst
}

func Assignment3() {
	// Create slice of teams
	teams := []Team{
		{"A", 5, 1, 2, 23, 20},
		{"B", 6, 1, 1, 24, 20},
		{"C", 7, 1, 0, 24, 10},
		{"D", 6, 1, 1, 15, 11},
		{"E", 8, 0, 1, 23, 12},
	}

	// Sort the teams based on score, goal difference, and goal for
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].Score() != teams[j].Score() {
			return teams[i].Score() > teams[j].Score()
		}

		if teams[i].GoalDiff() != teams[j].GoalDiff() {
			return teams[i].GoalDiff() > teams[j].GoalDiff()
		}

		return teams[i].GoalFor > teams[j].GoalFor
	})

	// Print the ranking
	fmt.Println("Team\tGoal For\tGoal Diff\tScore")
	for _, team := range teams {
		fmt.Printf("%s\t%d\t\t%d\t\t%d\n", team.Name, team.GoalFor, team.GoalDiff(), team.Score())
	}
}
