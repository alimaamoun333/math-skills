# Math Skills

A command-line statistical computation tool built in Go that reads numerical data from a file and calculates essential statistical metrics: **Average**, **Median**, **Variance**, and **Standard Deviation**, rounding each result to the nearest integer.

---

## 📋 Features

- **Average (Arithmetic Mean)**: Calculates the mean of all data points.
- **Median**: Computes the central tendency using an in-house sorting algorithm (`bubbleSort`).
- **Variance**: Measures the spread of data relative to the mean.
- **Standard Deviation**: Computes the measure of dispersion using a custom Newton-Raphson square root algorithm (`squareRoot`).
- **Custom Rounding**: Rounds results to the nearest integer (handles positive and negative numbers).
- **Robust Input Parsing**: Reads data line-by-line and validates float parsing.

---

## 📁 Project Structure

```text
math-skills/
├── data.txt                # Sample dataset
├── go.mod                  # Go module definition
├── main.go                 # Application entry point and CLI orchestration
├── README.md               # Project documentation
└── internal/
    ├── data/
    │   ├── reader.go       # File reading and parsing logic
    │   └── reader_test.go  # Unit tests for data reader
    └── stats/
        ├── stats.go        # Statistical functions and algorithms
        └── stats_test.go   # Comprehensive unit tests for stats calculations
```

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) (version 1.20 or later recommended)

### Usage

Run the program with a path to a file containing numbers (one per line):

```bash
go run . data.txt
```

### Example

Given `data.txt` containing:
```text
10
12
-15
20
25
0
-20
-2
100
```

Running the tool:
```bash
go run . data.txt
```

Output:
```text
Average: 14
Median: 10
Variance: 1113
Standard Deviation: 33
```

---

## 🧪 Running Tests

To run the complete test suite across all packages:

```bash
go test -v ./...
```

To run tests with coverage:

```bash
go test -cover ./...
```

---

## 📐 Mathematical Formulas

| Metric | Formula |
| :--- | :--- |
| **Average** | $\mu = \frac{1}{N}\sum_{i=1}^{N} x_i$ |
| **Median** | $x_{(N+1)/2}$ (odd $N$) or $\frac{x_{N/2} + x_{N/2+1}}{2}$ (even $N$) after sorting |
| **Variance** | $\sigma^2 = \frac{1}{N}\sum_{i=1}^{N} (x_i - \mu)^2$ |
| **Standard Deviation** | $\sigma = \sqrt{\sigma^2}$ |
