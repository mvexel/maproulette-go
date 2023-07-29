package main

import (
	"fmt"
	"log"

	"sr.ht/~mvexel/maproulette"
)

func main() {
	mr := maproulette.MapRoulette{APIKey: "yourapikey"}
	challenges, err := mr.GetChallenges()
	if err != nil {
		log.Fatalf("Error getting challenges: %v", err)
	}
	for _, challenge := range challenges {
		fmt.Printf("Challenge ID: %d, Name: %s\n", challenge.ID, challenge.Name)
	}
}
