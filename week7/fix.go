package main

func quickSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	pivot := a[0]
	var left, right []int
	for _, v := range a[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}
