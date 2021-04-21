package main

import "fmt"

func main() {

	fmt.Printf(largestTimeFromDigits([]int{0,0,2,1}))

}

func largestTimeFromDigits(arr []int) string {
	var hh,mm, j, iu, ju  int

	for i, val := range arr {

		fmt.Printf("i:%d, val:%d\n", i, val)
		for j = 0; j < 4; j++ {
			var h int
			if i == j {
				continue
			}

			fmt.Printf("j:%d\n", j)
			h = (val * 10) + arr[j]
			if (hh <= h) && (h < 24) {
				hh = h
				iu = i
				ju = j
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("%d:%d\n\n", hh, mm)

	for i, val := range arr {

		if iu == i {
			continue
		}

		fmt.Printf("i:%d, val:%d\n", i, val)
		for j = 0; j < 4; j++ {
			var m int
			if i == j || ju == j || iu == j {
				continue
			}

			fmt.Printf("j:%d\n", j)
			m = (val * 10) + arr[j]
			if (mm <= m) && (m < 60) {
				mm = m
			}
		}
		fmt.Printf("\n")
	}

	return fmt.Sprintf("%d:%d", hh, mm)
}
