package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	list1, list2 := getFileLists("input.txt")
	fmt.Println(getListsDistance(list1, list2))
}

func getFileLists(fileName string) ([]float64, []float64) {
	list1 := []float64{}
	list2 := []float64{}

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return []float64{}, []float64{}
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		columns := strings.Fields(line)

		firstValue, _ := strconv.Atoi(columns[0])
		secondValue, _ := strconv.Atoi(columns[1])

		list1 = append(list1, float64(firstValue))
		list2 = append(list2, float64(secondValue))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return list1, list2
}

func getListsDistance(list1 []float64, list2 []float64) int {
	sort.Float64s(list1)
	sort.Float64s(list2)

	totalDistance := 0.0

	for i := 0; i < len(list2); i++ {
		var difference float64 = list1[i] - list2[i]
		totalDistance += math.Abs(difference)
	}

	return int(totalDistance)
}
