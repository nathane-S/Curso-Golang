package main

import (
	"fmt"
	"strings"
)

func teamResult(team string, match []string) int {
	matchResult := match[2]
	teamfirst := match[0]

	if matchResult == "win" && teamfirst == team || teamfirst != team && matchResult == "loss" {
		return 3
	}
	if matchResult == "win" && teamfirst != team || teamfirst == team && matchResult == "loss" {
		return 0
	}
	return 1
}

func generateResult(match []string) map[string]int {
	teams := []string{match[0], match[1]}
	teamsResults := make(map[string]int, 0)
	for _, team := range teams {
		teamsResults[team] = teamResult(team, match)
	}
	return teamsResults
}

func main() {
	entrada := `Allegoric Alaskans;Blithering Badgers;win
Devastating Donkeys;Courageous Californians;draw`
	matches := strings.Split(entrada, "\n")

	fmt.Println(matches)
	matchResults := []map[string]int{}
	for _, match := range matches {
		arr := strings.Split(match, ";")
		matchResult := generateResult(arr)
		matchResults = append(matchResults, matchResult)
	}
	fmt.Println(matchResults)

	//Result := generateResult(arr)
	//fmt.Println(Result["Allegoric Alaskans"])

}
