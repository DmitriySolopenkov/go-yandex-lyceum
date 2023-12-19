package main

import (
	"os"
	"testing"
	"time"
)

func TestExtractLog(t *testing.T) {
	type test struct {
		fileContent string
		fileName    string
		outContent  []string
		wantError   bool
		start       time.Time
		end         time.Time
	}

	tests := []test{
		{
			fileContent: `12.12.2022 info
13.12.2022 info
14.12.2022 info
15.12.2022 info
16.12.2022 info`,
			start: time.Date(2022, 12, 14, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2022, 12, 15, 0, 0, 0, 0, time.UTC),

			fileName: "file1.txt",
			outContent: []string{
				"14.12.2022 info",
				"15.12.2022 info",
			},
		},
		{
			fileContent: `12.12.2022 info
13.12.2022 info
14.12.2022 info
15.12.2022 info
16.12.2022 info`,
			start: time.Date(2022, 12, 19, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2022, 12, 20, 0, 0, 0, 0, time.UTC),

			fileName:   "file2.txt",
			outContent: []string{},
			wantError:  true,
		},
		{
			fileContent: `12.12.2022 info
13.12.2022 info
14.12.2022 info
15.12.2022 info
16.12.2022 info`,
			fileName:   "file3.txt",
			outContent: []string{},
			wantError:  true,
		},
	}

	for _, tt := range tests {
		err := os.WriteFile(tt.fileName, []byte(tt.fileContent), 0644)
		if err != nil {
			t.Fatalf("Ощибка записи в файл %q: %s", tt.fileName, err)
		}

		defer os.Remove(tt.fileName)

		ExtractLog(tt.fileName, tt.start, tt.end)

	}
}
