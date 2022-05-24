package main

import "fmt"

//https://www.algoexpert.io/questions/Tournament%20Winner

func main() {
	fmt.Println(TournamentWinner([][]string{{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"}}, []int{0, 0, 1}))
}

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	currentBestTeam := ""
	teamPointsMap := make(map[string]int)
	teamPointsMap[currentBestTeam] = 0
	for index, wonTeam := range results {
		won := competitions[index][0]
		if wonTeam != HOME_TEAM_WON {
			won = competitions[index][1]
		}
		teamPointsMap[won] += 3
		if teamPointsMap[won] > teamPointsMap[currentBestTeam] {
			currentBestTeam = won
		}
	}
	return currentBestTeam
}
