package main

// import "fmt"

type calcInfo struct {
	maxSum, freeSum, childSum, werbeKostenLump, multiplier, socialMulti float64
}

type incomeInfo struct {
	workIncome, familyIncome, socialIncome, otherIncome, assets, balance float64
	children                                                             int
}

func main() {
	drawWindow()
}

func getMonthlyRate(income incomeInfo, calc calcInfo) float64 {
	biI := getBiannualIncome(income)
	moI := getMonthlyIncome(biI, calc.werbeKostenLump)
	social := moI * calc.socialMulti
	free := calc.freeSum + (calc.childSum * float64(income.children))
	return moI - social - free
}

func getFundingRate(income incomeInfo, calc calcInfo) float64 {
	monthlyRate := getMonthlyRate(income, calc)
	return (calc.maxSum - monthlyRate) * calc.multiplier
}

func getMonthlyIncome(biannualIncome, wkLump float64) float64 {
	return (biannualIncome - wkLump) / 6
}

func getBiannualIncome(income incomeInfo) float64 {
	biannualIncome := income.workIncome + income.familyIncome + income.socialIncome + income.otherIncome
	biannualIncome = biannualIncome * 6.
	return biannualIncome + income.assets + income.balance
}
