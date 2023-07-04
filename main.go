package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	ID     int
	Dices  []int
	Points int
}

func main() {
	var numPlayers, numDices int

	fmt.Print("Pemain= ")
	fmt.Scan(&numPlayers)

	fmt.Print("Dadu= ")
	fmt.Scan(&numDices)

	players := initializePlayers(numPlayers, numDices)
	fmt.Println("==================")

	round := 1
	for len(players) > 1 {
		fmt.Printf("Giliran %d lempar dadu:\n", round)

		for i := range players {
			rollDices(&players[i])
			fmt.Printf("Pemain #%d (%d): %v\n", players[i].ID, players[i].Points, players[i].Dices)
		}

		fmt.Println("Setelah evaluasi:")
		players = evaluateDices(players)

		for i := range players {
			fmt.Printf("Pemain #%d (%d): %v\n", players[i].ID, players[i].Points, players[i].Dices)
		}

		fmt.Println("==================")
		round++
	}

	fmt.Println("Game berakhir karena hanya pemain #", players[0].ID, "yang memiliki dadu.")
	fmt.Println("Game dimenangkan oleh pemain #", players[0].ID, "karena memiliki poin lebih banyak dari pemain lainnya.")
}

func initializePlayers(numPlayers, numDices int) []Player {
	players := make([]Player, numPlayers)

	for i := 0; i < numPlayers; i++ {
		players[i].ID = i + 1
		players[i].Dices = make([]int, numDices)
	}

	return players
}

func rollDices(player *Player) {
	rand.Seed(time.Now().UnixNano())

	for i := range player.Dices {
		player.Dices[i] = rand.Intn(6) + 1
	}
}

func evaluateDices(players []Player) []Player {
	for i := range players {
		player := &players[i]

		for j := 0; j < len(player.Dices); j++ {
			dice := player.Dices[j]

			switch dice {
			case 1:
				if i == len(players)-1 {
					players[0].Dices = append(players[0].Dices, dice)
				} else {
					players[i+1].Dices = append(players[i+1].Dices, dice)
				}

			case 6:
				player.Points++
			}
		}

		player.Dices = filterDices(player.Dices)
	}

	return filterPlayers(players)
}

func filterDices(dices []int) []int {
	filtered := []int{}

	for _, dice := range dices {
		if dice != 1 && dice != 6 {
			filtered = append(filtered, dice)
		}
	}

	return filtered
}

func filterPlayers(players []Player) []Player {
	filtered := []Player{}

	for _, player := range players {
		if len(player.Dices) > 0 {
			filtered = append(filtered, player)
		}
	}

	return filtered
}
