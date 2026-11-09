package stats

import (
	"math"
	"testing"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		name        string
		input       []float64
		expected    float64
		expectError bool
	}{
		{
			name:        "valid positive numbers",
			input:       []float64{1, 2, 3, 4, 5},
			expected:    3.0,
			expectError: false,
		},
		{
			name:        "single element",
			input:       []float64{42},
			expected:    42.0,
			expectError: false,
		},
		{
			name:        "floating point numbers",
			input:       []float64{1.5, 2.5, 3.5},
			expected:    2.5,
			expectError: false,
		},
		{
			name:        "numbers including negative values",
			input:       []float64{-10, 10, -5, 5},
			expected:    0.0,
			expectError: false,
		},
		{
			name:        "empty slice",
			input:       []float64{},
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil slice",
			input:       nil,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Average(tt.input)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error for input %v, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error for input %v: %v", tt.input, err)
			}

			tolerance := 1e-9
			if math.Abs(got-tt.expected) > tolerance {
				t.Errorf("Average(%v) = %v, expected %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestMedian(t *testing.T) {
	tests := []struct {
		name        string
		input       []float64
		expected    float64
		expectError bool
	}{
		{
			name:        "odd count unsorted",
			input:       []float64{5, 2, 1, 4, 3},
			expected:    3.0,
			expectError: false,
		},
		{
			name:        "even count unsorted",
			input:       []float64{10, 2, 8, 4},
			expected:    6.0,
			expectError: false,
		},
		{
			name:        "single element",
			input:       []float64{7.0},
			expected:    7.0,
			expectError: false,
		},
		{
			name:        "two elements",
			input:       []float64{1.5, 3.5},
			expected:    2.5,
			expectError: false,
		},
		{
			name:        "negative numbers",
			input:       []float64{-5, -1, -3},
			expected:    -3.0,
			expectError: false,
		},
		{
			name:        "empty slice",
			input:       []float64{},
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil slice",
			input:       nil,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy to check original slice is not mutated
			var original []float64
			if tt.input != nil {
				original = append([]float64(nil), tt.input...)
			}

			got, err := Median(tt.input)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error for input %v, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error for input %v: %v", tt.input, err)
			}

			tolerance := 1e-9
			if math.Abs(got-tt.expected) > tolerance {
				t.Errorf("Median(%v) = %v, expected %v", tt.input, got, tt.expected)
			}

			// Verify input was not modified
			if tt.input != nil {
				for i := range original {
					if tt.input[i] != original[i] {
						t.Errorf("Median mutated the input slice: got %v, original was %v", tt.input, original)
						break
					}
				}
			}
		})
	}
}

func TestVariance(t *testing.T) {
	tests := []struct {
		name        string
		input       []float64
		expected    float64
		expectError bool
	}{
		{
			name:        "standard integers",
			input:       []float64{1, 2, 3, 4, 5},
			expected:    2.0,
			expectError: false,
		},
		{
			name:        "single element",
			input:       []float64{10.0},
			expected:    0.0,
			expectError: false,
		},
		{
			name:        "identical elements",
			input:       []float64{4.0, 4.0, 4.0, 4.0},
			expected:    0.0,
			expectError: false,
		},
		{
			name:        "floating point numbers",
			input:       []float64{1.5, 2.5, 3.5, 4.5},
			expected:    1.25,
			expectError: false,
		},
		{
			name:        "negative numbers",
			input:       []float64{-4, -2, 2, 4},
			expected:    10.0,
			expectError: false,
		},
		{
			name:        "empty slice",
			input:       []float64{},
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil slice",
			input:       nil,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Variance(tt.input)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error for input %v, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error for input %v: %v", tt.input, err)
			}

			tolerance := 1e-9
			if math.Abs(got-tt.expected) > tolerance {
				t.Errorf("Variance(%v) = %v, expected %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestStandardDeviation(t *testing.T) {
	tests := []struct {
		name        string
		input       []float64
		expected    float64
		expectError bool
	}{
		{
			name:        "standard integers with exact integer SD",
			input:       []float64{2, 4, 4, 4, 5, 5, 7, 9},
			expected:    2.0,
			expectError: false,
		},
		{
			name:        "standard sequence",
			input:       []float64{1, 2, 3, 4, 5},
			expected:    math.Sqrt(2.0),
			expectError: false,
		},
		{
			name:        "single element",
			input:       []float64{10.0},
			expected:    0.0,
			expectError: false,
		},
		{
			name:        "identical elements",
			input:       []float64{4.0, 4.0, 4.0, 4.0},
			expected:    0.0,
			expectError: false,
		},
		{
			name:        "negative and positive numbers",
			input:       []float64{-4, -2, 2, 4},
			expected:    math.Sqrt(10.0),
			expectError: false,
		},
		{
			name:        "empty slice",
			input:       []float64{},
			expected:    0,
			expectError: true,
		},
		{
			name:        "nil slice",
			input:       nil,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StandardDeviation(tt.input)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error for input %v, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error for input %v: %v", tt.input, err)
			}

			tolerance := 1e-9
			if math.Abs(got-tt.expected) > tolerance {
				t.Errorf("StandardDeviation(%v) = %v, expected %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected int
	}{
		{"round positive down", 4.4, 4},
		{"round positive half up", 4.5, 5},
		{"round positive up", 4.6, 5},
		{"zero", 0.0, 0},
		{"round negative down", -4.4, -4},
		{"round negative half down", -4.5, -5},
		{"round negative up", -4.6, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Round(tt.input)
			if got != tt.expected {
				t.Errorf("Round(%v) = %v, expected %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected []float64
	}{
		{
			name:     "unsorted numbers",
			input:    []float64{5, 2, 8, 1, 9},
			expected: []float64{1, 2, 5, 8, 9},
		},
		{
			name:     "already sorted",
			input:    []float64{1, 2, 3, 4},
			expected: []float64{1, 2, 3, 4},
		},
		{
			name:     "reverse sorted",
			input:    []float64{4, 3, 2, 1},
			expected: []float64{1, 2, 3, 4},
		},
		{
			name:     "empty slice",
			input:    []float64{},
			expected: []float64{},
		},
		{
			name:     "single element",
			input:    []float64{42},
			expected: []float64{42},
		},
		{
			name:     "negative and duplicate values",
			input:    []float64{3, -1, 2, -1, 0},
			expected: []float64{-1, -1, 0, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]float64(nil), tt.input...)
			bubbleSort(nums)
			if len(nums) != len(tt.expected) {
				t.Fatalf("bubbleSort(%v) length = %v, expected %v", tt.input, len(nums), len(tt.expected))
			}
			for i := range nums {
				if nums[i] != tt.expected[i] {
					t.Errorf("bubbleSort(%v)[%d] = %v, expected %v", tt.input, i, nums[i], tt.expected[i])
				}
			}
		})
	}
}

func TestSquareRoot(t *testing.T) {
	tests := []struct {
		name        string
		input       float64
		expected    float64
		expectError bool
	}{
		{
			name:        "perfect square",
			input:       4.0,
			expected:    2.0,
			expectError: false,
		},
		{
			name:        "another perfect square",
			input:       9.0,
			expected:    3.0,
			expectError: false,
		},
		{
			name:        "zero",
			input:       0.0,
			expected:    0.0,
			expectError: false,
		},
		{
			name:        "non-integer square root",
			input:       2.0,
			expected:    math.Sqrt(2.0),
			expectError: false,
		},
		{
			name:        "negative number",
			input:       -4.0,
			expected:    0.0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := squareRoot(tt.input)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error for squareRoot(%v), got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error for squareRoot(%v): %v", tt.input, err)
			}

			tolerance := 1e-7
			if math.Abs(got-tt.expected) > tolerance {
				t.Errorf("squareRoot(%v) = %v, expected %v", tt.input, got, tt.expected)
			}
		})
	}
}



