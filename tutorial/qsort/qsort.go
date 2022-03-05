package qsort

func qsort(arr []int, left int, right int) {
	var i int = left
	var j int = right
	var pivot int = arr[left]
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
