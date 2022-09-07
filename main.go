package main

import (
	"fmt"
        "os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(calcDMI())
}

func readFile(filename string) []string {
	dat, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	split := strings.Split(string(dat), "\n")
	return split
}

// calcDMI calculates the depth measurement increase for input. That is, the value between two values is positive or negative. Count the number of positive values (values where the current number is greater than the previous number).
func calcDMI() int {
        testInput := readFile("/home/jmorris/Dev/go-workspace/src/adventOfCode2021/problems/problem1.csv")
	n := len(testInput)
        a := testInput[1:n]
	b := testInput[0:n-1]

        var tmp []int
	for i := 0; i < len(a); i++ {
		x, _ := strconv.Atoi(a[i])
		y, _ := strconv.Atoi(b[i])
		tmp = append(tmp, x-y)
	}

	var count int
	for i := 0; i < len(tmp); i++ {
		if tmp[i] > 0 {
			count++
		}
	}
        return count
}
