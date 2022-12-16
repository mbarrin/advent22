package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	rowID int
	colID int
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	beacons := map[point]bool{}
	sensors := map[point]bool{}
	invalidPoints := map[point]bool{}

	for scanner.Scan() {
		line := scanner.Text()
		var sensorRow, sensorCol, beaconRow, beaconCol int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorRow, &sensorCol, &beaconRow, &beaconCol)

		beacon := point{rowID: beaconRow, colID: beaconCol}
		sensor := point{rowID: sensorRow, colID: sensorCol}

		distance := manhattenDistance(beacon.colID, sensor.colID, beacon.rowID, sensor.rowID)

		invalid(sensor, invalidPoints, distance)

		beacons[beacon] = true
		sensors[sensor] = true
	}
	fmt.Println("part 1:", invalidCount(invalidPoints, 2000000))
}

func manhattenDistance(x1, x2, y1, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func invalidCount(points map[point]bool, row int) int {
	count := 0

	for k := range points {
		if k.rowID == row {
			count++
		}
	}

	return count
}

func invalid(sensor point, points map[point]bool, distance int) {
	for row := 0; row < distance; row++ {
		for col := 0; col < distance; col++ {
			if manhattenDistance(sensor.colID, sensor.colID+col, sensor.rowID, sensor.rowID+row) <= distance {
				points[point{rowID: sensor.rowID + row, colID: sensor.colID + col}] = true
				points[point{rowID: sensor.rowID - row, colID: sensor.colID + col}] = true
				points[point{rowID: sensor.rowID + row, colID: sensor.colID - col}] = true
				points[point{rowID: sensor.rowID - row, colID: sensor.colID - col}] = true
			}
		}
	}
}
