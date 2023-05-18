package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the input file as a command-line argument")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}
	defer file.Close()

	player1Wins := 0
	player2Wins := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		hand := strings.Fields(line)

		player1Hand := hand[:5]
		player2Hand := hand[5:]

		player1Rank := calculateRank(player1Hand)
		player2Rank := calculateRank(player2Hand)

		if player1Rank > player2Rank {
			player1Wins++
		} else if player1Rank < player2Rank {
			player2Wins++
		}
	}

	fmt.Printf("Player 1: %d hands\n", player1Wins)
	fmt.Printf("Player 2: %d hands\n", player2Wins)
}

func calculateRank(hand []string) int {
	rank := 0

	if isPair(hand) {
		rank = 1
	} else if isTwoPairs(hand) {
		rank = 2
	} else if isThreeOfAKind(hand) {
		rank = 3
	} else if isFlush(hand) {
		rank = 4
	} else if isFullHouse(hand) {
		rank = 5
	}

	return rank
}

func isPair(hand []string) bool {
	cardCounts := make(map[string]int)

	for _, card := range hand {
		value := card[:len(card)-1]
		cardCounts[value]++
	}

	for _, count := range cardCounts {
		if count == 2 {
			return true
		}
	}

	return false
}

func isTwoPairs(hand []string) bool {
	cardCounts := make(map[string]int)
	pairCount := 0

	for _, card := range hand {
		value := card[:len(card)-1]
		cardCounts[value]++
	}

	for _, count := range cardCounts {
		if count == 2 {
			pairCount++
		}
	}

	return pairCount == 2
}

func isThreeOfAKind(hand []string) bool {
	cardCounts := make(map[string]int)

	for _, card := range hand {
		value := card[:len(card)-1]
		cardCounts[value]++
	}

	for _, count := range cardCounts {
		if count == 3 {
			return true
		}
	}

	return false
}

func isFlush(hand []string) bool {
	suit := hand[0][len(hand[0])-1]

	for i := 1; i < len(hand); i++ {
		if hand[i][len(hand[i])-1] != suit {
			return false
		}
	}

	return true
}

func isFullHouse(hand []string) bool {
	cardCounts := make(map[string]int)
	hasThree := false
	hasPair := false

	for _, card := range hand {
		value := card[:len(card)-1]
		cardCounts[value]++
	}

	for _, count := range cardCounts {
		if count == 3 {
			hasThree = true
		} else if count == 2 {
			hasPair = true
		}
	}

	return hasThree && hasPair
}
