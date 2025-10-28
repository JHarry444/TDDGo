
package shipping

func CalculateShippingCost(orderAmount int) int {
	if orderAmount < 0 {
		return -1 // invalid input
	}
	if orderAmount >= 1000 {
		return 0
	}
	return 50
}
