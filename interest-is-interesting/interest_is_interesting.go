package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return 3.213
	case balance >= 1000 && balance < 5000:
		return 1.621
	case balance >= 5000:
		return 2.475
	default:
		return 0.5
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	interest := float64(InterestRate(balance)) / 100
	return balance * interest
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	interest := float64(InterestRate(balance)) / 100
	return balance * (1.0 + interest)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	i := 0
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		i++
	}
	return i
}
