package Test

import (
	Algo "DSA/Algorithms"
	"reflect"
	"testing"
)

func TestSolveMaze(t *testing.T) {
	tests := []struct {
		name     string
		maze     [][]string
		wall     string
		start    []int
		end      []int
		expected bool
	}{
		{
			name: "Basic path",
			maze: [][]string{
				{"x", " ", "x"},
				{"x", " ", "x"},
				{"x", " ", "x"},
			},
			wall:     "x",
			start:    []int{0, 1},
			end:      []int{2, 1},
			expected: true,
		},
		{
			name: "Blocked path",
			maze: [][]string{
				{"x", " ", "x"},
				{"x", "x", "x"},
				{"x", " ", "x"},
			},
			wall:     "x",
			start:    []int{0, 1},
			end:      []int{2, 1},
			expected: false,
		},
		{
			name: "Start is end",
			maze: [][]string{
				{"x", " ", "x"},
			},
			wall:     "x",
			start:    []int{0, 1},
			end:      []int{0, 1},
			expected: true,
		},
		{
			name: "Big maze with path",
			maze: [][]string{
				{"x", "x", "x", "x", "x", "x", "x", "x", "x", "x", " ", "x"},
				{"x", " ", " ", " ", " ", " ", " ", " ", " ", "x", " ", "x"},
				{"x", " ", " ", " ", " ", " ", " ", " ", " ", "x", " ", "x"},
				{"x", " ", "x", "x", "x", "x", "x", "x", "x", "x", " ", "x"},
				{"x", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "x"},
				{"x", " ", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x"},
			},
			wall:     "x",
			start:    []int{0, 10},
			end:      []int{5, 1},
			expected: true,
		},
		{
			name: "No wall maze (open field)",
			maze: [][]string{
				{" ", " ", " "},
				{" ", " ", " "},
				{" ", " ", " "},
			},
			wall:     "x",
			start:    []int{0, 0},
			end:      []int{2, 2},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := Algo.SolveMaze(tt.maze, tt.wall, tt.start, tt.end)
			got := len(path) > 0 && reflect.DeepEqual(path[len(path)-1], tt.end)
			if got != tt.expected {
				t.Errorf("expected path found: %v, got: %v, path: %v", tt.expected, got, path)
			}
		})
	}
}
