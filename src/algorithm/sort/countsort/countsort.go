package countsort

type CountSort struct {
	Value []int
}

func (c *CountSort) Sort() ([]int, error) {
	max, min := c.Value[0], c.Value[0]
	for _, v := range c.Value {
		if v > max {
			max = v
		}
		if min > v {
			min = v
		}
	}
	temp := make([]int, max-min+1)
	for _, v := range c.Value {
		temp[v-min]++
	}
	for i := 1; i < max-min+1; i++ {
		temp[i] += temp[i-1]
	}
	result := make([]int, len(c.Value))
	for i := 0; i < len(c.Value); i++ {
		temp[c.Value[i]-min]--
		result[temp[c.Value[i]-min]] = c.Value[i]
	}
	c.Value = result
	return result, nil
}
