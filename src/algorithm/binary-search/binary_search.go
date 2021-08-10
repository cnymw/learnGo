package binary_search

// 二分查找递归实现
func binarySearchRecursive(arr []int, l, r, x int) int {
	if r >= l {
		mid := l + (r-l)/2

		if arr[mid] == x {
			return mid
		}

		if arr[mid] > x {
			return binarySearchRecursive(arr, l, mid-1, x)
		}

		return binarySearchRecursive(arr, mid+1, r, x)
	}
	return -1
}

// 二分查找迭代实现
func binarySearchIterative(arr []int, l, r, x int) int {
	for l <= r {
		mid := l + (r-l)/2

		if arr[mid] == x {
			return mid
		}

		if arr[mid] < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
