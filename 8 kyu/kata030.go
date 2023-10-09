package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	numStrs := strings.Split(strings.ReplaceAll(strings.TrimSpace(str), " ", ""), ",")
	nums := make([]int, len(numStrs))
	for i, str := range numStrs {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Number conversion error:", err)
			return
		}
		nums[i] = num
	}
	fmt.Println(SmallestIntegerFinder(nums))
}
func SmallestIntegerFinder(numbers []int) int {
	minNum := math.MaxInt64
	for _, n := range numbers {
		if n < minNum {
			minNum = n
		}
	}
	return minNum
}
