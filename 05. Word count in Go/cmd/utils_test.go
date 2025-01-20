package cmd

import (
	"reflect"
	"testing"
)

func TestGetTextFileCounts(t *testing.T) {

	tests := []struct {
		fileName       string
		expectedCounts TextFileCounts
		wantErr        bool
	}{
		{"test_1.txt", TextFileCounts{lineCount: 6, wordCount: 98, characterCount: 642}, false},
		{"test_2.txt", TextFileCounts{lineCount: 12, wordCount: 151, characterCount: 1007}, false},
		{"bar", TextFileCounts{lineCount: 0, wordCount: 0, characterCount: 0}, true},
		{"not_exists.txt", TextFileCounts{lineCount: 0, wordCount: 0, characterCount: 0}, true},
	}

	for _, test := range tests {

		receivedCounts, err := getTextFileCounts(test.fileName)

		if (err != nil) != test.wantErr {
			t.Errorf("getTextFileCounts() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		if !reflect.DeepEqual(receivedCounts, test.expectedCounts) {
			t.Errorf("getTextFileCounts() = %v, want %v", receivedCounts, test.expectedCounts)
		}

	}

}
