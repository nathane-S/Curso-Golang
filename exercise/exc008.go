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
	var teamsResults map[string]int
	for _, team := range teams {
	}
}

func main() {
	entrada := `Allegoric Alaskans;Blithering Badgers;draw`
	arr := strings.Split(entrada, ";")
	fmt.Println(arr[2])

	Result := teamResult(arr[1], arr)
	fmt.Println(Result)
}
