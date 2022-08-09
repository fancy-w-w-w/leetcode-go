package specialquestion

// 盛水最多容器
func MaxArea(height []int) int {
	leftPoint, rightPoint := 0, len(height)-1
	res := 0
	for leftPoint < rightPoint {
		value, isLeft := minNum(height[leftPoint], height[rightPoint])
		tmpWater := value * (rightPoint - leftPoint)
		if tmpWater > res {
			res = tmpWater
		}
		// 移动指针
		if isLeft {
			leftPoint++
		} else {
			rightPoint--
		}
	}
	return res
}

func minNum(a int, b int) (int, bool) {
	if a < b {
		return a, true
	}
	return b, false
}
