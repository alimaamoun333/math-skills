package stats

import (
	"fmt"
)

func Average(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no number found ")
	}
	sum := 0.0
	for _, number := range numbers {
		sum += number

	}
	return sum / float64(len(numbers)), nil
}

func Median(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no number found ")
	}
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)
	bubbleSort(sorted)
	if len(sorted)%2 != 0 {
		return sorted[len(sorted)/2], nil
	}
	firstIndex := (len(sorted) / 2) - 1
	secondIndex := len(sorted) / 2
	result, err := Average(sorted[firstIndex : secondIndex+1])
	return result, err
}

func Variance(numbers []float64) (float64, error) {
	mean, err := Average(numbers)
	if err != nil {
		return 0, fmt.Errorf("can't find the mean %w", err)
	}
	sum := 0.0
	for _, number := range numbers {
		difference := number - mean
		squaredDifference := difference * difference
		sum += squaredDifference
	}
	return sum / float64(len(numbers)), nil
}

func StandardDeviation(numbers []float64) (float64, error) {
	value, err := Variance(numbers)
	if err != nil {
		return 0, fmt.Errorf("calculate standard deviation: %w", err)
	}
	result, err := squareRoot(value)
	if err != nil {
		return 0, fmt.Errorf("can'r get square %w", err)
	}
	return result, nil
}

func bubbleSort(numbers []float64) {
	n := len(numbers)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

}

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("cannot calculate square root of negative number")
	}
	if number == 0 {
		return 0, nil
	}
	guess := number / 2
	tolerance := 1e-9
	for {

		newGuess := (guess + number/guess) / 2

		difference := guess - newGuess

		if difference < 0 {
			difference = -difference
		}

		if difference < tolerance {
			return newGuess, nil
		}
		guess = newGuess
	}
}

func Round(number float64) int {
	if number >= 0 {
		return int(number + 0.5)
	}
	return int(number - 0.5)
}
