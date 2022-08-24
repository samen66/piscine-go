package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	} else {
		result := make([]int, max-min)
		var j int = 0
		for i := min; i < max; i++ {
			result[j] = i
			j++
		}
		return result
	}
}
