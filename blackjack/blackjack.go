package blackjack

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
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	switch hand := ParseCard(card1) + ParseCard(card2); {
	case hand < 21 && hand >= 17:
		return "S"
	case hand == 22:
		return "P"
	case hand == 21:
		switch ParseCard(dealerCard) < 10 {
		case true:
			return "W"
		default:
			return "S"
		}
	case hand < 17 && hand >= 12:
		switch ParseCard(dealerCard) >= 7 {
		case true:
			return "H"
		default:
			return "S"
		}
	default:
		return "H"
	}
}
