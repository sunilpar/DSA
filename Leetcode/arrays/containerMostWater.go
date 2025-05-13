package leetcode

func MaxArea(height []int) int {
	max := 0
	for i, iv := range height {
		diff := 0
		sm := 0
		for j, jv := range height {
			if j > i {
				diff = j - i
			} else {
				diff = i - j
			}
			if jv > iv {
				sm = iv
			} else {
				sm = jv
			}
			area := sm * diff
			if area > max {
				max = area
			}
		}
	}
	return max
}
