// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func bubsort(list []int) []int {
	var tmp int
	for i := 1; i < len(list); i++ {
		for j := 1; j < len(list); j++ {
			if list[j-1] > list[j] {
				tmp = list[j]
				list[j] = list[j-1]
				list[j-1] = tmp
			}
		}
	}
	return list
}
func merge(trick []int, trit []int) []int {
	len1 := len(trick)
	len2 := len(trit)
	var lolo []int
	trick = bubsort(trick)
	trit = bubsort(trit)
	ick := 0
	it := 0
	for i := 0; i < len1+len2; i++ {
		if ick == len1 {
			lolo = append(lolo, trit[it])
			it++
		} else if it == len2 {
			lolo = append(lolo, trick[ick])
			ick++
		} else {
			if trick[ick] < trit[it] {
				lolo = append(lolo, trick[ick])
				ick++
			} else {
				lolo = append(lolo, trit[it])
				it++
			}
		}
	}

	return lolo
}
func main() {
	first := []int{3, 8, 2, 9}
	second := []int{8, 54, 4, 67, 1}

	fmt.Println(bubsort(first))
	fmt.Println(merge(first, second))

}

