package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Result is: %v == [2]int{6, 766}", NbMonths(2000, 8000, 1000, 1.5))
}

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	net_worth := float64(startPriceOld)
	price_old := float64(startPriceOld)
	price_new := float64(startPriceNew)
	percentage := percentLossByMonth
	months := 0
	savings := 0.0

	for months = 0; net_worth < price_new; months++ {
		price_old = (100 - percentage) / 100 * price_old
		price_new = (100 - percentage) / 100 * price_new
		if months%2 == 0 {
			percentage += 0.5
		}
		savings += float64(savingperMonth)
		net_worth = price_old + savings
		fmt.Printf("End of the month %v\n", months)
		fmt.Printf("Net worth: %v\n", net_worth)
		fmt.Printf("New car price: %v\n", price_new)
	}

	return [2]int{months, int(math.Round(net_worth - price_new))}
}
