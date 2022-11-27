package hashmap

func twoSum(nums []int, target int) []int {
	// hashTable 通过 map 结构记录结果
	hashTable := map[int]int{}

	for i, x := range nums {
		// 从 hashTable 里直接取对应的结果。
		// 例如当前为 x=2，target 为 9，那么取 hashTable[7]
		// 如果取的到话，直接返回结果
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		// 取不到的话，存入hashTable
		// 例如当前 x=2，存入 hashTable[2]
		hashTable[x] = i
	}
	return []int{0, 0}
}
