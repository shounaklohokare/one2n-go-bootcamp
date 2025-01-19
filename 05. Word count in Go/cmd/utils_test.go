package cmd

import (
	"testing"
)

func TestGetLineCount(t *testing.T) {

	tests := []struct {
		fileName          string
		expectedLineCount int
		wantErr           bool
	}{
		{"test_1.txt", 6, false},
		{"test_2.txt", 12, false},
		{"bar", -1, true},
		{"not_exists.txt", -1, true},
	}

	for _, test := range tests {

		gotLineCount, err := getLineCount(test.fileName)

		if (err != nil) != test.wantErr {
			t.Errorf("getLineCount() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		if gotLineCount != test.expectedLineCount {
			t.Errorf("getLineCount() = %v, want %v", gotLineCount, test.expectedLineCount)
		}

	}

}

func TestGetWordCount(t *testing.T) {

	tests := []struct {
		fileName          string
		expectedWordCount int
		wantErr           bool
	}{
		{"test_1.txt", 98, false},
		{"test_2.txt", 151, false},
		{"bar", -1, true},
		{"not_exists.txt", -1, true},
	}

	for _, test := range tests {

		gotWordCount, err := getWordCount(test.fileName)

		if (err != nil) != test.wantErr {
			t.Errorf("getWordCount() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		if gotWordCount != test.expectedWordCount {
			t.Errorf("getWordCount() = %v, want %v", gotWordCount, test.expectedWordCount)
		}

	}

}
