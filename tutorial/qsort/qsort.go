package qsort

func qsort(arr []int, left int, right int) {
	var i int = left
	var j int = right
	// select middle one as a pivot
	var chosen int = left + ((right - left) >> 1)
	var pivot int = arr[chosen]
	// exchange it to the first element
	arr[left], arr[chosen] = arr[chosen], arr[left]
	for i < j {
		for i < j && arr[j] >= pivot {
			j--
		}
		arr[i] = arr[j]
		for i < j && arr[i] < pivot {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = pivot

	if left < i-1 {
		qsort(arr, left, i-1)
	}
	if i+1 < right {
		qsort(arr, i+1, right)
	}
}
