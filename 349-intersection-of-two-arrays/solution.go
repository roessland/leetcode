package main

func intersection(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]bool)
	for _, n := range nums1 {
		set1[n] = true
	}

	set2 := make(map[int]bool)
	for _, n := range nums2 {
		set2[n] = true
	}

	intersection := []int{}
	for n := range set1 {
		if set2[n] {
			intersection = append(intersection, n)
		}
	}

	return intersection
}

func main() {

}
