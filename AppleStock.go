package main

import "fmt"

/**
 * Apple Stock Main.
 *
 * This is a Go solution for the Apple Stock technical interview
 * question.
 */
func main() {
	stockPricesYesterday := []int{10, 7, 5, 8, 11, 9}
	fmt.Print(getMaxProfit(stockPricesYesterday))
}

/**
 * The maximum profit can be found by keeping the lowest
 * value found so far, and comparing it to each of the values
 * at the current index to check the maximum profit that can be
 * obtained by selling it at each index.
 */
func getMaxProfit(stockPricesYesterday[] int) (int) {

	lowestBuyPrice := stockPricesYesterday[0];
	maxProfit := 0;
	for i := 0; i < len(stockPricesYesterday); i++ {
		currentSalePrice := stockPricesYesterday[i]
		var currentProfit int = currentSalePrice - lowestBuyPrice
		if currentProfit > maxProfit {
			maxProfit = currentProfit;
			lowestBuyPrice = currentSalePrice
		}
		if currentSalePrice < lowestBuyPrice {
			lowestBuyPrice = currentSalePrice
		}
	}
	return maxProfit
}