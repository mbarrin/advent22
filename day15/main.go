package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type point struct {
	rowID int
	colID int
}

var lowestCol, highestCol = math.Inf(1), math.Inf(-1)

func main() {
	file, _ := os.Open("input.txt")
	//file, _ := os.Open("sample.txt")
	scanner := bufio.NewScanner(file)

	beacons := map[point]bool{}
	sensors := map[point]int{}

	for scanner.Scan() {
		line := scanner.Text()
		var sensorRow, sensorCol, beaconRow, beaconCol int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorCol, &sensorRow, &beaconCol, &beaconRow)

		beacon := point{rowID: beaconRow, colID: beaconCol}
		sensor := point{rowID: sensorRow, colID: sensorCol}

		distance := manhattenDistance(beacon.colID, sensor.colID, beacon.rowID, sensor.rowID)

		if float64(beaconCol-distance) < lowestCol {
			lowestCol = float64(beaconCol - distance)
		}

		if float64(beaconCol+distance) > highestCol {
			highestCol = float64(beaconCol + distance)
		}

		sensors[sensor] = distance
		beacons[beacon] = true
	}
	//fmt.Println("part 1:", invalidCount(lowestCol, highestCol, 10, sensors, beacons))
	fmt.Println("part 1:", invalidCount(lowestCol, highestCol, 2000000, sensors, beacons))
}

func invalidCount(start, end float64, row int, sensors map[point]int, beacons map[point]bool) int {
	count := 0
	matches := map[point]bool{}

	for col := start; col <= end; col++ {
		for p, d := range sensors {
			distance := manhattenDistance(int(col), p.colID, row, p.rowID)
			if distance <= d {
				tmp := point{row, int(col)}
				_, sensorExists := sensors[tmp]
				if !matches[tmp] && !beacons[tmp] && !sensorExists {
					matches[tmp] = true
					count++
				}
			}
		}
	}
	return count
}

func manhattenDistance(col1, col2, row1, row2 int) int {
	return abs(col1-col2) + abs(row1-row2)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
