package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king", "queen", "jack", "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardOnHand := ParseCard(card1) + ParseCard(card2)
	switch {
	case cardOnHand == 22:
		return "P"
	case cardOnHand == 21 && ParseCard(dealerCard) < 10:
		return "W"
	case cardOnHand == 21:
		return "S"
	case cardOnHand >= 17 && cardOnHand <= 20:
		return "S"
	case (cardOnHand >= 12 && cardOnHand <= 16) && ParseCard(dealerCard) >= 7:
		return "H"
	case (cardOnHand >= 12 && cardOnHand <= 16):
		return "S"
	default:
		return "H"
	}
}
