package selectsort

// Sort list of int using select algorithm
func Sort(list []int) []int {

	for i := 0; i < len(list)-1; i++ {
		idx := i
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[idx] {
				idx = j
			}
		}

		temp := list[i]
		list[i] = list[idx]
		list[idx] = temp
	}

	return list
}
