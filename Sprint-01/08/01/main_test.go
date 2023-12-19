package main

import (
	"os"
	"testing"
)

func TestSaveState(t *testing.T) {
	type test struct {
		cells     [][]bool
		fileName  string
		content   string
		wantError bool
	}

	tests := []test{
		{
			cells: [][]bool{
				{true, true},
				{false, false},
			},
			fileName: "file1",
			content:  "11\n00",
		},

		{
			cells: [][]bool{
				{true, true, false},
				{false, false, true},
			},
			fileName: "file2",
			content:  "110\n001",
		},
		{
			cells: [][]bool{
				{true, true},
				{false, false},
			},
			fileName:  "",
			content:   "",
			wantError: true,
		},
	}
	for _, tt := range tests {
		defer os.Remove(tt.fileName)
		w := World{
			Height: len(tt.cells),
			Width:  len(tt.cells[0]),
			Cells:  tt.cells,
		}
		err := w.SaveState(tt.fileName)

		if err != nil && tt.wantError {
			continue
		}

		if err != nil {
			t.Fatalf("TestSaveState failed with an error: %s", err)
		}

		if tt.wantError {
			t.Fatalf("TestSaveState ожидаемый an error here")
		}

		fileContent, err := os.ReadFile(tt.fileName)
		if err != nil {
			t.Fatalf("TestSaveState failed to read the file: %s", err)
		}

		content := string(fileContent)

		if content != tt.content {
			t.Fatalf("TestSaveState failed, ожидаемый content:\n%s\nbut got:\n%s", tt.content, content)
		}
	}
}
