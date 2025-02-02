package entity

type Cart struct {
	Items    []Item
	Discount Discount
}

func (c *Cart) GetTotalPrice() float64 {
	totalPrice := 0.0
	for _, item := range c.Items {
		totalPrice += item.Price
	}

	totalPrice = applyDiscount(totalPrice, c.Discount.Percentage)

	return totalPrice
}

func applyDiscount(totalPrice, discountPercentage float64) float64 {
	return totalPrice * (1 - discountPercentage)
}
