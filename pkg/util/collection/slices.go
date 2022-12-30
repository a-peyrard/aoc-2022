package collection

func Sum(slice *[]int) int {
	result := 0
	for _, elem := range *slice {
		result += elem
	}
	return result
}
