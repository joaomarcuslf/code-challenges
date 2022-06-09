package challenge04

import (
	"sort"
)

func ExtractList(k []int) (int, int) {
	sort.Ints(k)

	repeated := 0
	for i := 0; i < len(k); i++ {
		if i == len(k)-1 {
			continue
		}

		if k[i] == k[i+1] {
			repeated++

			if i+1 < len(k) {
				i++
			}
		}
	}

	unique := len(k) - repeated*2

	return unique, repeated
}
