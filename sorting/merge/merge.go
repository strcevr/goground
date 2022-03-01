package merge

import "sync"

func Sort(data []int) []int {
	return mergeSort(data)
}

func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	var l []int
	var r []int

	m := len(data) / 2

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(){
		defer wg.Done()
		l = mergeSort(data[:m])
	}()
	go func(){
		defer wg.Done()
		r = mergeSort(data[m:])
	}()
	wg.Wait()

	return mergeSortCombine(l, r)
}

func mergeSortCombine(l, r []int) []int {
	c := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		switch {
		case len(l) == 0:
			return append(c, r...)
		case len(r) == 0:
			return append(c, l...)
		case l[0] <= r[0]:
			c = append(c, l[0])
			l = l[1:]
		case l[0] >= r[0]:
			c = append(c, r[0])
			r = r[1:]
		}
	}
	return c
}
