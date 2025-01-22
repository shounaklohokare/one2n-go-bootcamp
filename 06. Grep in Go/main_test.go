package main

import (
	"reflect"
	"testing"
)

func TestGetStringOccurencesInAFile(t *testing.T) {

	tests := []struct {
		searchString   string
		fileName       string
		expectedOutput []string
		wantErr        bool
	}{
		{"DevOps", "test_1.txt", []string{"The adoption of DevSecOps emphasizes incorporating security measures throughout the DevOps pipeline.", "Artificial intelligence and machine learning are increasingly integrated into DevOps to automate tasks and enhance predictive analytics."}, false},
	}

	for _, test := range tests {

		receivedOutput, err := getStringOccurencesInAFile(test.searchString, test.fileName)

		if (err != nil) != test.wantErr {
			t.Errorf("getStringOccurencesInAFile() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		if !reflect.DeepEqual(receivedOutput, test.expectedOutput) {
			t.Errorf("getStringOccurencesInAFile() = %v, want %v", receivedOutput, test.expectedOutput)
		}

	}

}
