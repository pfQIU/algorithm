package test

import (
	"algorithm/code"
	"testing"
)

func TestStock(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	if code.MaxProfit(prices) != 5 {
		t.Errorf("maxProfit(prices) = %d; want 5", code.MaxProfit(prices))
	}
}
