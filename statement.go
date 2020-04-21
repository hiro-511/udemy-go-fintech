package main

import (
	"fmt"
)

// Q1 . 以下のスライスから一番小さい数を探して出力するコードを書いてください。

// l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

// Q2. 以下の果物の価格の合計を出力するコードを書いてください。

// m := map[string]int{
//     "apple":  200,
//     "banana": 300,
//     "grapes": 150,
//     "orange": 80,
//     "papaya": 500,
//     "kiwi":   90,
// }

func main() {
	// Q1
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	fmt.Println("Minimum Number ->", getMinNum(l))

	// Q2
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	fmt.Println("Total Number ->", getTotalNum(m))
}

func getMinNum(l []int) int {
	var min int
	for i, num := range l {
		if i == 0 {
			min = num
			continue
		}

		if min >= num {
			min = num
		}
	}
	return min
}

func getTotalNum(m map[string]int) int {
	total := 0
	for _, v := range m {
		total = total + v
	}
	return total
}
