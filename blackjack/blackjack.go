// package blackjack
package main

import (
	"fmt"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

/*cards value
| card  | value | card    | value |
| :---: | :---: | :-----: | :---: |
|  ace  |  11   | eight   |   8   |
|  two  |   2   | nine    |   9   |
| three |   3   |  ten    |  10   |
| four  |   4   | jack    |  10   |
| five  |   5   | queen   |  10   |
|  six  |   6   | king    |  10   |
| seven |   7   | *other* |   0   |*/
// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
/*- If you have a pair of aces you must always split them.
- If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace,
a figure or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand
and wait for the reveal of the other card.
- If your cards sum up to a value within the range [17, 20] you should always stand.
- If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher,
 in which case you should always hit.
- If your cards sum up to 11 or lower you should always hit. "S-Stand", "P-Split", "H-Hit"
*/
func FirstTurn(card1, card2, dealerCard string) string {
	var card1_number_value, card2_number_value, dealer_card_number_value int = ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)
	var sum_of_cards int = card1_number_value + card2_number_value
	switch {
	case sum_of_cards == 22:
		return "P"
	case sum_of_cards > 21 && dealer_card_number_value < 10:
		return "W"
	case sum_of_cards > 21 && dealer_card_number_value >= 10:
		return "S"
	case sum_of_cards >= 17 && sum_of_cards <= 20:
		return "S"
	case (sum_of_cards >= 12 && sum_of_cards <= 16) && dealer_card_number_value <= 7:
		return "S"
	case (sum_of_cards >= 12 && sum_of_cards <= 16) && dealer_card_number_value > 7:
		return "H"
	case sum_of_cards <= 11:
		return "H"
	default:
		return "H"
	}

}

func main() {
	fmt.Println(FirstTurn("ace", "ten", "ace"))
}
