package main

import "sort"

func findContentChildren(greeds []int, sizes []int) int {
	sortedGreeds := make([]int, len(greeds))
	copy(sortedGreeds, greeds)
	sort.Ints(sortedGreeds)

	sortedSizes := make([]int, len(sizes))
	copy(sortedSizes, sizes)
	sort.Ints(sortedSizes)

	contentChildren := 0
	for g, s := len(sortedGreeds)-1, len(sortedSizes)-1; g >= 0 && s >= 0; g-- {
		if sortedGreeds[g] <= sortedSizes[s] {
			contentChildren++
			s--
		}
	}
	return contentChildren
}

func main() {

}
