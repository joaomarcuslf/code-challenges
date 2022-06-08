package challenge04

import (
	"sort"
)

func ExtractList(k []int) (int, int) {
	sort.Ints(k)

	unique := 0
	repeated := 0
	for i := 0; i < len(k); i++ {
		if i == len(k)-1 {
			unique++
			continue
		} else if k[i] == k[i+1] {
			repeated++

			if i+1 < len(k) {
				i++
			}
		} else {
			unique++
		}
	}
	return unique, repeated
}
