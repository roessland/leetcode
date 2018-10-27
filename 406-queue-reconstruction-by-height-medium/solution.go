package leetcode406

import "sort"

/**
(7,0),(7,1),(6,1),(5,0),(5,2),(4,4)

(5,0)(5,2)(7,0)(6,1)(4,4)(7,1)
*/

func sortPeople(people [][]int) {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] >= people[j][0]
	})
}

func reconstructQueue(people [][]int) [][]int {
	sortPeople(people)
	q := [][]int{}
	for _, p := range people {
		q = append(q, []int{0, 0})
		copy(q[p[1]+1:], q[p[1]:])
		q[p[1]] = p
	}
	return q
}
