package main

import "testing"

func TestCopyFilePart(t *testing.T) {
	type args struct {
		inputFilename string
		outFileName   string
		startpos      int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CopyFilePart(tt.args.inputFilename, tt.args.outFileName, tt.args.startpos); (err != nil) != tt.wantErr {
				t.Errorf("CopyFilePart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
