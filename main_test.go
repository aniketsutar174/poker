package main

import "testing"

func TestCalculateRank(t *testing.T) {
	tests := []struct {
		hand         []string
		expectedRank int
	}{
		{[]string{"AH", "AS", "AD", "AC", "2H"}, 5}, // Full House
		{[]string{"3D", "6D", "7D", "TD", "QD"}, 4}, // Flush
		{[]string{"2C", "2D", "3H", "3S", "3C"}, 3}, // Three of a Kind
		{[]string{"2C", "2D", "3H", "3S", "4C"}, 2}, // Two Pairs
		{[]string{"2C", "2D", "3H", "4S", "5C"}, 1}, // Pair
		{[]string{"2C", "3D", "4H", "5S", "7C"}, 0}, // High Card
	}

	for _, test := range tests {
		rank := calculateRank(test.hand)
		if rank != test.expectedRank {
			t.Errorf("Hand: %v, Expected Rank: %d, Got: %d", test.hand, test.expectedRank, rank)
		}
	}
}
