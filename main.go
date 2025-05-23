package main

import (
	"GolangBackendAssessment/core"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// run process log file
	var filePath string
	fmt.Print("Enter FilePath: ")
	fmt.Scan(&filePath)

	fmt.Print("Enter keywords (separated by comma): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	keywords := strings.Split(scanner.Text(), ",")

	keywordCounts, err := core.ProcessLogFile(filePath, keywords)
	if err != nil {
		fmt.Println("Error Processing log file:", err)
		return
	}
	for _, v := range keywordCounts {
		log.Println(v)
	}

	// run prime panlidrome
	fmt.Println()
	fmt.Println("Question two")
	var N int
	fmt.Print("Enter N: ")
	fmt.Scan(&N)

	primePalindromes := core.FindPrimePalindromes(N)
	fmt.Printf("Sum of the first %d prime palindromic numbers: %d\n", N, primePalindromes)
}
