package math

import (
	"sort"
)

func CalculateMean(arr []int) float64 {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	mean := float64(sum) / float64(len(arr))
	return mean
}

func CalculateMedian(arr []int) float64 {
	sort.Ints(arr)

	length := len(arr)
	if length%2 == 0 {
		midIndex := length / 2
		return float64(arr[midIndex-1]+arr[midIndex]) / 2
	} else {
		midIndex := length / 2
		return float64(arr[midIndex])
	}
}

func CalculateMode(arr []int) []int {
	freqMap := make(map[int]int)

	// 統計頻率
	for _, num := range arr {
		freqMap[num]++
	}

	// 找到最高數字
	maxFreq := 0
	for _, freq := range freqMap {
		if freq > maxFreq {
			maxFreq = freq
		}
	}

	// 收集所有最高數字
	mode := []int{}
	for num, freq := range freqMap {
		if freq == maxFreq {
			mode = append(mode, num)
		}
	}

	return mode
}