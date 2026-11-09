package main

import (
	"fmt"
	"math-skills/internal/data"
	"math-skills/internal/stats"
	"os"
)

func main() {
	err := Run()
	if err != nil {
		fmt.Println(err)
	}
}
func Run() error {
	args := os.Args
	if len(args) != 2 {
		return fmt.Errorf("THE USEAGE OF THE PROGRAM : go run . data.txt")
	}
	file := args[1]
	nums, err := data.Read(file)
	if err != nil {
		return fmt.Errorf("read data: %w", err)
	}
	average, err := stats.Average(nums)
	if err != nil {
		return fmt.Errorf("calculate average: %w", err)
	}
	median, err := stats.Median(nums)
	if err != nil {
		return fmt.Errorf("calculate median: %w", err)
	}
	variance, err := stats.Variance(nums)
	if err != nil {
		return fmt.Errorf(" calculate variance: %w", err)
	}
	standardDeviation, err := stats.StandardDeviation(nums)
	if err != nil {
		return fmt.Errorf("calculate standard deviation: %w", err)
	}
	fmt.Println("Average:", stats.Round(average))
	fmt.Println("Median:", stats.Round(median))
	fmt.Println("Variance:", stats.Round(variance))
	fmt.Println("Standard Deviation:", stats.Round(standardDeviation))
	return nil

}
