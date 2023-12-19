package main

import (
	"os"
	"testing"
)

func TestReadContent(t *testing.T) {
	type test struct {
		fileContent string
		fileName    string
		wantError   bool
	}

	tests := []test{
		{
			fileContent: `Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.`,
			fileName:    "file1.txt",
			wantError:   false,
		},
		{
			fileContent: `
                Golang outperforms Python in terms of microservices, 
                APIs, and other fast-loading feature`,
			fileName:  "file2.txt",
			wantError: false,
		},
		{
			fileContent: `
                Golang outperforms Python in terms of microservices, 
                APIs, and other fast-loading feature`,
			fileName:  "file3.txt",
			wantError: true,
		},
	}

	for _, tt := range tests {

		err := os.WriteFile(tt.fileName, []byte(tt.fileContent), 0644)
		if err != nil {
			t.Fatalf("Fail to write a file %q: %s", tt.fileName, err)
		}

		if tt.wantError {
			os.Remove(tt.fileName)
		}
	}
}
