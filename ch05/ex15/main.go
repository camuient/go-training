package main

import (
	"fmt"
	"math"
)

func max(vals ...int) (int, error) {
	m := math.MinInt32

	if len(vals) == 0 {
		return m, fmt.Errorf("引数がありません")
	}

	for _, val := range vals {
		if val > m {
			m = val
		}
	}
	return m, nil
}

func min(vals ...int) (int, error) {
	m := math.MaxInt32

	if len(vals) == 0 {
		return m, fmt.Errorf("引数がありません")
	}

	for _, val := range vals {
		if val < m {
			m = val
		}
	}
	return m, nil
}

func max2(v int, vals ...int) int {
	m := v
	for _, val := range vals {
		if val > m {
			m = val
		}
	}
	return m
}

func min2(v int, vals ...int) int {
	m := v
	for _, val := range vals {
		if val < m {
			m = val
		}
	}
	return m
}
