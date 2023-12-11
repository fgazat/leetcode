package besttimetobuystock

func maxProfit(prices []int) int {
	var (
		minPrice  = prices[0]
		maxProfix int
	)
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfix {
			maxProfix = price - minPrice
		}
	}
	return maxProfix
}
