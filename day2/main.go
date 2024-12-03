package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	exampleInput := getDataSetFromFile("input.txt")
	scoreMap := map[string]int{"safe": 0, "unsafe": 0}

	for i := 0; i < len(exampleInput); i++ {
		if isSafe(exampleInput[i]) {
			scoreMap["safe"]++
		} else {
			scoreMap["unsafe"]++
		}
	}
	fmt.Println(scoreMap)
}

func isSafe(sequence []int) bool {

	for i := 1; i < len(sequence); i++ {
		operation := "Increase"

		if sequence[1] < sequence[0] {
			operation = "Decrease"
		}

		difference := sequence[i] - sequence[i-1]

		if difference == 0 ||
			(operation == "Increase" && (difference > 3 || difference < 0)) ||
			(operation == "Decrease" && (difference < -3 || difference > 0)) {
			return false
		}

	}
	return true
}

func getDataSetFromFile(fileName string) [][]int {

	dataSet := [][]int{}

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("ERROR OPENING FILE")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		columns := strings.Fields((line))

		newArray := []int{}

		for i := 0; i < len(columns); i++ {
			num, err := strconv.Atoi(columns[i])

			if err != nil {
				fmt.Println("Error converting string to int:", err)
			}
			newArray = append(newArray, num)
		}

		dataSet = append(dataSet, newArray)
	}

	return dataSet

}
