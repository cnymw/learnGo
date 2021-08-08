package p922

func sortArrayByParityII(A []int) []int {
	for i := 0; i < len(A); i++ {
		if i%2 == 0 && A[i]%2 != 0 {
			for j := i + 1; j < len(A); j++ {
				if A[j]%2 == 0 {
					A[i], A[j] = A[j], A[i]
					break
				}
			}
		} else if i%2 != 0 && A[i]%2 == 0 {
			for j := i + 1; j < len(A); j++ {
				if A[j]%2 != 0 {
					A[i], A[j] = A[j], A[i]
					break
				}
			}
		}
	}
	return A
}
