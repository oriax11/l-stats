package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numbers []int
	if len(os.Args) != 2 {
		fmt.Println("Please Provide Data correctly")
		os.Exit(0)
	}
	File, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(0)
	}
	Filecontent := string(File)
	Filecontent = strings.ReplaceAll(Filecontent, "\r\n", "\n")

	numbstring := strings.Split(Filecontent, "\n")
	for i := range numbstring {
		numbstring[i] = strings.TrimSpace(numbstring[i])
	}
	var num int
	for i := range numbstring {
		if numbstring[i] == "" {
			continue
		}
		num, err = strconv.Atoi(numbstring[i])
		if err != nil {
			fmt.Println(err)
			fmt.Println("So Please Provide a valid data format")
			os.Exit(0)
		}
		numbers = append(numbers, num)
	}

	// Calculate Linear Regression Line
	m, b := linearRegression(numbers)
	fmt.Printf("Linear Regression Line: y = %.2fx + %.2f\n", m, b)

	// Calculate Pearson Correlation Coefficient
	r := pearsonCorrelation(numbers)
	fmt.Printf("Pearson Correlation Coefficient: %.2f\n", r)
}

func linearRegression(y []int) (float64, float64) {
	n := float64(len(y))
	var sumX, sumY, sumXY, sumX2 float64

	for i, yi := range y {
		x := float64(i + 1)
		sumX += x
		sumY += float64(yi)
		sumXY += x * float64(yi)
		sumX2 += x * x
	}

	// Calculate slope (m) and intercept (b)
	m := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	b := (sumY - m*sumX) / n

	return m, b
}

func pearsonCorrelation(y []int) float64 {
	n := float64(len(y))
	var sumX, sumY, sumXY, sumX2, sumY2 float64

	for i, yi := range y {
		x := float64(i + 1)
		sumX += x
		sumY += float64(yi)
		sumXY += x * float64(yi)
		sumX2 += x * x
		sumY2 += float64(yi) * float64(yi)
	}

	// Calculate Pearson Correlation Coefficient
	numerator := n*sumXY - sumX*sumY
	denominator := math.Sqrt((n*sumX2 - sumX*sumX) * (n*sumY2 - sumY*sumY))

	return numerator / denominator
}
