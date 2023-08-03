package main

import (
	"fmt"
	"math/rand"
)

type player struct {
	ID        int
	Dice      []int
	Points    int
	NextPlayer *player
}

func createPlayers(numPlayers, numDice int) []*player {
	players := make([]*player, numPlayers)
	for i := 0; i < numPlayers; i++ {
		players[i] = &player{
			ID:   i + 1,
			Dice: make([]int, numDice),
		}
	}

	for i := 0; i < numPlayers-1; i++ {
		players[i].NextPlayer = players[i+1]
	}
	players[numPlayers-1].NextPlayer = players[0]

	return players
}

func rollDice(player *player) {
	for i := range player.Dice {
		player.Dice[i] = rand.Intn(6) + 1
	}
}

func evaluateDice(player *player) {
	remainingDice := make([]int, 0)
	nextPlayer := player.NextPlayer

	for _, dice := range player.Dice {
		switch dice {
		case 6:
			player.Points++
		case 1:
			nextPlayer.Dice = append(nextPlayer.Dice, 1)
		default:
			remainingDice = append(remainingDice, dice)
		}
	}

	player.Dice = remainingDice
}

func findWinner(players []*player) *player {
	winner := players[0]
	for _, player := range players[1:] {
		if player.Points > winner.Points {
			winner = player
		}
	}
	return winner
}

func playGame(players []*player) {
	round := 1
	activePlayers := len(players)

	for activePlayers > 1 {
		fmt.Printf("==================\n")
		fmt.Printf("Giliran %d lempar dadu:\n", round)

		for _, player := range players {
			if len(player.Dice) > 0 {
				rollDice(player)
				fmt.Printf("Pemain #%d (%d): %v\n", player.ID, player.Points, player.Dice)
			}
		}

		for _, player := range players {
			if len(player.Dice) > 0 {
				evaluateDice(player)
			}
		}

		activePlayers = 0
		for _, player := range players {
			if len(player.Dice) > 0 {
				activePlayers++
			}
		}

		round++
	}

	fmt.Printf("==================\n")
	fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu.\n", players[0].ID)
	winner := findWinner(players)
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.\n", winner.ID)
}


func main() {
	totalPlayers := 3
	totalDice := 4

	players := createPlayers(totalPlayers, totalDice)
	playGame(players)
}
