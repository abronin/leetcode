package main

func main() {
}

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	filtered := constraints(restaurants, veganFriendly, maxPrice, maxDistance)
	if len(filtered) == 0 {
		return []int{}
	}
	res := orderIDs(restaurants, filtered)
	return res
}

func orderIDs(restaurants [][]int, filtered []int) []int {
	n := len(filtered)
	res := []int{filtered[0]}

	for i := 1; i < n; i++ {
		rest := restaurants[filtered[i]]
		id, rating := rest[0], rest[1]
		index := -1
		for j := 0; j < i; j++ {
			if index == -1 {
				jrest := restaurants[res[j]]
				jrating := jrest[1]
				if rating > jrating {
					index = j
				} else if rating == jrating {
					kid, krating := id+1, rating
					for k := j; kid > id && rating == krating && k < i; k++ {
						krest := restaurants[res[k]]
						kid, krating = krest[0], krest[1]
						index = k
					}
				}
			}
		}
		if index == -1 {
			index = i
		}
		res = insert(res, index, filtered[i])
	}

	for i, index := range res {
		res[i] = restaurants[index][0]
	}

	return res
}

func constraints(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	filtered := []int{}
	for i, rest := range restaurants {
		if veganFriendly == 1 && rest[2] != 1 {
			continue
		}
		if maxPrice < rest[3] {
			continue
		}
		if maxDistance < rest[4] {
			continue
		}
		filtered = append(filtered, i)
	}
	return filtered
}

func insert(arr []int, index, value int) []int {
	res := make([]int, len(arr)+1)
	for i := 0; i < len(res); i++ {
		if i < index {
			res[i] = arr[i]
		} else if i == index {
			res[i] = value
		} else {
			res[i] = arr[i-1]
		}
	}
	return res
}
