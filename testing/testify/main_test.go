package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type product struct {
	Product_id int `json:"product_id"`
	Name string `json:"name"`
	Price int `json:"price"`
	In_stock bool `json:"in_stock"`
}

func TestIsNonNegative(t *testing.T) {
	cases := []struct{
		name string
		product product
	}{
		{
			name: "Happy Path",
			product: product{
				Product_id: 1,
				Name: "Banana",
				Price: 10,
				In_stock: true,
			},
		},
		{
			name: "Small Negative",
			product: product{
				Product_id: 2,
				Name: "Phone",
				Price: -1000,
				In_stock: false,
			},
		},
		{
			name: "Big Negative 1",
			product: product{
				Product_id: 3,
				Name: "Tomato",
				Price: -2,
				In_stock: true,
			},
		},
		{
			name: "Big Negative 2",
			product: product{
				Product_id: 3,
				Name: "Potato",
				Price: -1,
				In_stock: true,
			},
		},
		{
			name: "Zero Price",
			product: product{
				Product_id: 4,
				Name: "Seeds",
				Price: 0,
				In_stock: true,
			},
		},
		{
			name: "Small Postive 1",
			product: product{
				Product_id: 5,
				Name: "Beer",
				Price: 1,
				In_stock: false,
			},
		},
		{
			name: "Small Positive 2",
			product: product{
				Product_id: 5,
				Name: "Wine",
				Price: 2,
				In_stock: true,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, tc.product.Price > 0)
		})
	}
}
