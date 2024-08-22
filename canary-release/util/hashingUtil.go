package util

func HashCode(s string) int {
	hash := 0
	for _, c := range s {
		hash = 31*hash + int(c)
	}
	if hash < 0 {
		hash = -hash
	}
	return hash
}

func HashRangeCounts(records [][]string) []int {
	hashRangeCounts := make([]int, 10) // Array to count ranges 0-9, 10-19, ..., 90-99

	for _, record := range records {
		if len(record) > 1 {
			deviceID := record[0]
			hashValue := HashCode(deviceID)
			canaryPercentage := (hashValue%100 + 100) % 100
			rangeIndex := canaryPercentage / 10
			if rangeIndex >= len(hashRangeCounts) {
				rangeIndex = len(hashRangeCounts) - 1
			}
			hashRangeCounts[rangeIndex]++
		}
	}

	return hashRangeCounts
}
