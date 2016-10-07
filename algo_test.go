package selectsort

import "testing"

var tests = []struct {
	list   []int
	result []int
}{
	{
		[]int{5, 1, 8, 4, 3, 7, 6, 2},
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
	},
	{
		[]int{58, 5, 7, 1, 12, 9},
		[]int{1, 5, 7, 9, 12, 58},
	},
}

func TestSearchSuccess(t *testing.T) {
	for _, test := range tests {
		res := Sort(test.list)
		if compareList(res, test.result) == false {
			t.Error("Error for list %v ,expected %v", res, test.result)
		}
	}
}

func compareList(la, lb []int) bool {
	if len(la) != len(lb) {
		return false
	}

	for i := 0; i < len(la); i++ {
		if la[i] != lb[i] {
			return false
		}
	}

	return true
}
