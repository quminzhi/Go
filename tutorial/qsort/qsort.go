package qsort

func qsort(arr []int, left int, right int) {
	if left < right {
		var i int = left
		var j int = right
		var pivot int = arr[left]
		for i < j {
			for i < right && arr[i] < pivot {
				i++
			}
			for j > left && arr[j] >= pivot {
				j--
			}
			if i < j {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		qsort(arr, left, j)
		qsort(arr, j+1, right)
	}
}
