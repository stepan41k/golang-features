package main

import "testing"

func TestCalculateDiscount(t *testing.T) {
	cases := []struct{
		name string
		price int
		expected int
	}{
		{
			name: "No  discount",
			price: 500,
			expected: 500,
		},
		{
			name: "10% discount",
			price: 1500,
			expected: 1350,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := CaclulateDiscount(tc.price)

			if res != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, res)
			}
		})
	}
}