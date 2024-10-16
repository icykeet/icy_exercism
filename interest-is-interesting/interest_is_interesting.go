package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return 3.213
	case balance >= 0 && balance < 1000:
		return 0.5
	case balance >= 1000 && balance < 5000:
		return 1.621
	default:
		return 2.475 // for value more than or equal 5000;
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	getRate := InterestRate(balance)
	precent := (float64(getRate / 100))
	return balance * precent
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	countYears := 0
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		countYears += 1
	}
	return countYears
}
