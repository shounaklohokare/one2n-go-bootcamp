package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetStringOccurencesInAFile(t *testing.T) {

	tests := []struct {
		searchString   string
		fileName       string
		expectedOutput []string
		expectedError  string
	}{
		{"DevOps", "test_1.txt", []string{"Artificial intelligence and machine learning are increasingly integrated into DevOps to automate tasks and enhance predictive analytics.",
			"The adoption of DevSecOps emphasizes incorporating security measures throughout the DevOps pipeline."}, ""},
		{"Docker", "not_exists.txt", nil, "No such file or directory"},
		{"Docker", "bar", nil, "Is a directory"},
	}

	for _, test := range tests {

		receivedOutput, err := getStringOccurencesInAFile(test.searchString, test.fileName)
		if err != nil {

			if !strings.Contains(err.Error(), test.expectedError) {
				t.Errorf("error = %v, wantErr %v", err, test.expectedError)
				return
			}

		}
		if !reflect.DeepEqual(receivedOutput, test.expectedOutput) {
			t.Errorf("got = %v, want %v", receivedOutput, test.expectedOutput)
		}

	}

}
