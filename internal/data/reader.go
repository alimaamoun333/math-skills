package data

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Read(path string) ([]float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	result := []float64{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("parse %q as number: %w", line, err)
		}
		result = append(result, number)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	return result, nil
}
