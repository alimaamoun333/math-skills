package data

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	t.Run("valid integers file", func(t *testing.T) {
		tempDir := t.TempDir()
		filePath := filepath.Join(tempDir, "data.txt")
		content := "189\n113\n121\n114\n145\n"

		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}

		got, err := Read(filePath)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []float64{189, 113, 121, 114, 145}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Read() = %v, expected %v", got, expected)
		}
	})

	t.Run("valid float and negative numbers", func(t *testing.T) {
		tempDir := t.TempDir()
		filePath := filepath.Join(tempDir, "data_floats.txt")
		content := "12.34\n-5.67\n0.0\n"

		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}

		got, err := Read(filePath)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []float64{12.34, -5.67, 0.0}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Read() = %v, expected %v", got, expected)
		}
	})

	t.Run("empty file", func(t *testing.T) {
		tempDir := t.TempDir()
		filePath := filepath.Join(tempDir, "empty.txt")

		if err := os.WriteFile(filePath, []byte(""), 0644); err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}

		got, err := Read(filePath)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(got) != 0 {
			t.Errorf("Read() returned non-empty slice %v for empty file", got)
		}
	})

	t.Run("invalid non-numeric line", func(t *testing.T) {
		tempDir := t.TempDir()
		filePath := filepath.Join(tempDir, "invalid.txt")
		content := "123\nabc\n456\n"

		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}

		got, err := Read(filePath)
		if err == nil {
			t.Fatalf("expected error for invalid data, got result: %v", got)
		}
	})

	t.Run("non-existent file", func(t *testing.T) {
		tempDir := t.TempDir()
		filePath := filepath.Join(tempDir, "missing.txt")

		got, err := Read(filePath)
		if err == nil {
			t.Fatalf("expected error for non-existent file, got result: %v", got)
		}
	})
}
