package main

const HOME_TEAM_WON = 1

/*
competitions = [["HTML", "C#"], ["C#", "Python"], ["Python", "C#"]]
results = [0, 0, 1]
output = "Python"
*/

func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.
	teamPoints := make(map[string]int)
	winnerOfTheTour := ""
	for i, value := range results {
		idx := 0
		if value == 0 {
			idx = 1
		}
		winnerTeam := competitions[i][idx]
		val := teamPoints[winnerTeam]
		teamPoints[winnerTeam] = val + 3
		if winnerOfTheTour == "" {
			winnerOfTheTour = winnerTeam
			continue
		}
		if teamPoints[winnerOfTheTour] < teamPoints[winnerTeam] {
			winnerOfTheTour = winnerTeam
		}
	}
	return winnerOfTheTour
}
