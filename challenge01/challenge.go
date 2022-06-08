package challenge01

// Merge Sort
func SortArrayMergeSort(arr []int) []int {
	var output []int

	// Use those index to store the index of the element on Output array
	var zeros []int
	var ones []int
	var twos []int

	for _, element := range arr {
		switch element {
		case 0:
			zeros = append(zeros, element)
		case 1:
			ones = append(ones, element)
		case 2:
			twos = append(twos, element)
		default:
		}
	}
	output = append(output, zeros...)
	output = append(output, ones...)
	output = append(output, twos...)

	return output
}

// Counting Sort (Somewhat)
func SortArrayCountingSort(arr []int) []int {
	var output []int

	qtZero := 0
	qtOne := 0
	qtTwo := 0

	for i := 0; i < len(arr); i++ {
		switch arr[i] {
		case 0:
			qtZero++
		case 1:
			qtOne++
		case 2:
			qtTwo++
		default:
		}
	}

	for i := 0; i < qtZero; i++ {
		output = append(output, 0)
	}
	for i := 0; i < qtOne; i++ {
		output = append(output, 1)
	}
	for i := 0; i < qtTwo; i++ {
		output = append(output, 2)
	}

	return output
}

// Quicksort
func SortArrayQuickSort(arr []int) []int {

	lo := 0
	mid := 0
	hi := len(arr) - 1

	temp := 0

	for mid <= hi {
		if arr[mid] == 0 {
			temp = arr[lo]

			arr[lo] = arr[mid]
			arr[mid] = temp

			lo++
			mid++
		} else if arr[mid] == 1 {
			mid++
		} else {
			temp = arr[mid]
			arr[mid] = arr[hi]
			arr[hi] = temp
			hi--
		}
	}

	return arr
}

func SortArrayBubbleSort(arr []int) []int {
	var output []int

	output = arr

	for i := 0; i < len(output)-1; i++ {
		for j := 0; j < len(output)-i-1; j++ {
			if output[j] > output[j+1] {
				output[j], output[j+1] = output[j+1], output[j]
			}
		}
	}

	return output
}
