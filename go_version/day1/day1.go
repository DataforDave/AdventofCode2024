package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {

	if x < 0 {
		return -x
	}
	return x
}

func CalculateSimilarity(slice0 []int, slice1 []int) (similaritySum int) {

	for i := 0; i < len(slice0); i++ {

		// Grab the value at index of slice0 compare to all other items in slice 1
		similarityMultiplier := 0
		for j := 0; j < len(slice1); j++ {

			if slice0[i] == slice1[j] {
				similarityMultiplier++

			}

		}

		similaritySum += slice0[i] * similarityMultiplier
		similarityMultiplier = 0

	}

	return
}

func GetElementDifferenceSum(slice0 []int, slice1 []int) (diffSlice []int) {

	for i := 0; i < len(slice0); i++ {

		diff := slice0[i] - slice1[i]
		diffSlice = append(diffSlice, abs(diff))

	}

	return
}

func ParseAndSort(fileLines []string) (slice0 []int, slice1 []int) {

	for _, line := range fileLines {
		stringSlice := strings.Fields(line)
		intVar0, err := strconv.Atoi(stringSlice[0])
		if err != nil {
			panic(err)
		}
		intVar1, err := strconv.Atoi(stringSlice[1])
		if err != nil {
			panic(err)
		}

		slice0 = append(slice0, intVar0)
		slice1 = append(slice1, intVar1)
	}
	sort.Ints(slice0)
	sort.Ints(slice1)

	return
}

func sumList(numbers []int) int {

	if len(numbers) == 0 {
		return 0
	}

	return numbers[0] + sumList(numbers[1:])
}

func main() {

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()
	slice0, slice1 := ParseAndSort(fileLines)
	diffSlice := GetElementDifferenceSum(slice0, slice1)
	fmt.Println(sumList(diffSlice))
	fmt.Println(CalculateSimilarity(slice0, slice1))

}
