package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := makeScanner()

	_ = readInt(scanner)
	k := readInt(scanner)
	list := readArray(scanner)

	fmt.Println(CalcMax(list, k))
}

func CalcMax(list []int, k int) int {
	var sum = sum(list[:k])
	var maxSum = sum
	var size = len(list)

	for i := 0; i < k; i++ {
		sum = sum - list[k-1-i] + list[size-1-i]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

func sum(numbs []int) int {
	result := 0
	for _, numb := range numbs {
		result += numb
	}

	return result
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
