package main

import (
	"os"
	"strings"
	"testing"
)

func TestModifyFile(t *testing.T) {
	type test struct {
		fileContent string
		fileName    string
		val         string
		startPos    int
		// wantError      bool
		outFileContent string
	}

	tests := []test{
		{
			fileContent:    `this text is modified`,
			fileName:       "file1.txt",
			val:            "body",
			outFileContent: "this body is modified",
			startPos:       5},
		{
			fileContent:    `some word to modify`,
			fileName:       "file2.txt",
			val:            "text",
			outFileContent: `some text to modify`,
			startPos:       5},
	}
	for _, tt := range tests {
		err := os.WriteFile(tt.fileName, []byte(tt.fileContent), 0644)
		if err != nil {
			t.Fatalf("Fail to write a file %q: %s", tt.fileName, err)
		}

		defer os.Remove(tt.fileName)

		ModifyFile(tt.fileName, tt.startPos, tt.val)

		content, err := os.ReadFile(tt.fileName)
		if err != nil {
			t.Fatalf("TestModifyFile failed on file %s with an error: %s", tt.fileName, err)
		}
		if !strings.Contains(string(content), tt.outFileContent) {
			t.Errorf("Expected content %q not found in the modified file %q", tt.outFileContent, tt.fileName)
		}
	}
}
