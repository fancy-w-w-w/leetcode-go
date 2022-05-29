package digit

func MajorityElement(nums []int) int {

	var amount int
	tmpStore := nums[0]
	for _, value := range nums {
		if tmpStore == value {
			amount++
		} else if amount > 1 {
			amount--
		} else {
			amount = 1
			tmpStore = value
		}
	}
	return tmpStore
}
