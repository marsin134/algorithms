package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shell(items []int, count int) {
	var i, j, gap, k int
	var x int
	a := [5]int{9, 5, 3, 2, 1}

	for k = 0; k < 5; k++ {
		gap = a[k]
		for i = gap; i < count; i++ {
			x = items[i]
			for j = i - gap; (x < items[j]) && (j >= 0); j = j - gap {
				items[j+gap] = items[j]
			}
			items[j+gap] = x
		}
	}
}

func qs(items []int, left int, right int) {
	var i, j int
	var x, y int

	i = left
	j = right
	x = items[(left+right)/2]

	for {
		for (items[i] < x) && (i < right) {
			i++
		}
		for (x < items[j]) && (j > left) {
			j--
		}

		if i <= j {
			y = items[i]
			items[i] = items[j]
			items[j] = y
			i++
			j--
		}

		if !(i <= j) {
			break
		}
	}

	if left < j {
		qs(items, left, j)
	}
	if i < right {
		qs(items, i, right)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const SIZE = 10000
	var a, b, c [SIZE][SIZE]int
	var elem_c int

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			a[i][j] = rand.Intn(100) + 1
		}
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			b[i][j] = rand.Intn(100) + 1
		}
	}

	start := time.Now()
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			elem_c = 0
			for r := 0; r < SIZE; r++ {
				elem_c = elem_c + a[i][r]*b[r][j]
				c[i][j] = elem_c
			}
		}
	}
	elapsed := time.Since(start)
	println("Size:", SIZE, "x", SIZE)
	fmt.Printf("Time taken for matrix multiplication: %.5f seconds\n", elapsed.Seconds())
}
