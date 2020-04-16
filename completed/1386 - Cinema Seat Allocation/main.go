package main

import "fmt"

func main() {
	fmt.Println("Run tests!")
}

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	busy := markBusy(reservedSeats)
	res := countBusy(busy)
	res += (n - len(busy)) * 2
	return res
}

func markBusy(reservedSeats [][]int) map[int][]bool {
	families := make(map[int][]bool)
	for _, reserve := range reservedSeats {
		row := reserve[0]
		seat := reserve[1]

		if families[row] == nil {
			families[row] = make([]bool, 3)
		}

		if seat == 2 || seat == 3 {
			families[row][0] = true
		}

		if seat == 8 || seat == 9 {
			families[row][2] = true
		}

		if seat == 4 || seat == 5 {
			families[row][0] = true
			families[row][1] = true
		}

		if seat == 6 || seat == 7 {
			families[row][1] = true
			families[row][2] = true
		}
	}
	return families
}

func countBusy(families map[int][]bool) int {
	res := 0
	for _, v := range families {
		if !v[0] && !v[1] && !v[2] {
			res += 2
			continue
		}
		if !v[0] && !v[1] && !v[2] {
			res++
			continue
		}
		if !v[0] && !v[1] && !v[2] {
			res++
			continue
		}
		if !v[0] && !v[1] && !v[2] {
			res++
			continue
		}
		if !v[0] || !v[1] || !v[2] {
			res++
		}
	}
	return res
}
