package main

import (
	"canary-release/util"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello Go, it's been a while, hasn't it? Let's get back to becoming GoGods")
	records, err := readCSV()
	if err != nil {
		fmt.Println("error opening file", err)
		return
	}

	err = writeDetailedCSV(records)
	if err != nil {
		fmt.Println("error reading the file ", err)
		return
	}

	hashRangeCounts := util.HashRangeCounts(records)

	err = writeRangeCSV(hashRangeCounts)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func readCSV() ([][]string, error) {
	file, err := os.Open("device-id.csv")
	if err != nil {
		return nil, fmt.Errorf("error opening file %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func writeDetailedCSV(records [][]string) error {
	file, err := os.Create("detailed-canary-release.csv")
	if err != nil {
		return fmt.Errorf("error in creating detailed release %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"device_id", "hash_code", "canary_percentage", "user_uid"})
	for _, record := range records {
		if len(records) > 0 {
			deviceId := record[0]
			hashCode := util.HashCode(deviceId)
			canaryPercentage := (hashCode%100 + 100) % 100
			userUid := record[1]
			writer.Write([]string{deviceId, fmt.Sprintf("%d", hashCode), fmt.Sprintf("%d", canaryPercentage), fmt.Sprint("%s", userUid)})
		}
	}
	return nil
}

func writeRangeCSV(hashRangeCounts []int) error {
	file, err := os.Create("range-canary-release.csv")
	if err != nil {
		return fmt.Errorf("error in creating range csv %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Range", "Count"})
	for i, count := range hashRangeCounts {
		rangeLabel := fmt.Sprintf("Below %d", (i+1)*10)
		writer.Write([]string{rangeLabel, fmt.Sprintf("%d", count)})
	}

	return nil
}
