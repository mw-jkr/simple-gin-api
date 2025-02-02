package entity_test

import (
	"example/test-golang/entity"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CartSuite struct {
	suite.Suite
	Cart  entity.Cart
	Items []entity.Item
}

func (s *CartSuite) SetupSuite() {
	item1 := entity.Item{Name: "item1", Price: 100.00}
	item2 := entity.Item{Name: "item2", Price: 200.00}
	item3 := entity.Item{Name: "item3", Price: 300.00}
	items := []entity.Item{item1, item2, item3}

	s.Items = items
	s.Cart = entity.Cart{Items: items}
}

func (s *CartSuite) TestGetTotalPrice() {
	assert := s.Assertions

	assert.Equal(600.00, s.Cart.GetTotalPrice())
}

func (s *CartSuite) TestGetTotalPriceWithDiscount() {
	assert := s.Assertions

	discount := entity.Discount{Percentage: 0.2}
	s.Cart.Discount = discount

	assert.Equal(480.0, s.Cart.GetTotalPrice())
}

func TestCartSuite(t *testing.T) {
	suite.Run(t, new(CartSuite))
}
