package main

import (
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestGrepFile(t *testing.T) {

	tests := []struct {
		searchString   string
		fileName       string
		expectedOutput []string
		expectedError  string
	}{
		{"DevOps", "test_1.txt", []string{"Artificial intelligence and machine learning are increasingly integrated into DevOps to automate tasks and enhance predictive analytics.",
			"The adoption of DevSecOps emphasizes incorporating security measures throughout the DevOps pipeline."}, ""},
		{"Docker", "not_exists.txt", nil, "invalid argument"},
		{"Docker", "bar", nil, "Incorrect function."},
	}

	for _, test := range tests {

		file, _ := os.Open(test.fileName)

		defer file.Close()

		receivedOutput, err := grep(file, test.searchString)
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

func TestGrepSTDIN(t *testing.T) {

	file, err := os.Open("test_1.txt")
	if err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin

	defer func() { os.Stdin = oldStdin }()

	os.Stdin = file
	var got []string
	if got, err = grep(file, "DevOps"); err != nil {
		t.Errorf("user input failed :- %v", err)
	}

	defer file.Close()
	want := []string{"Artificial intelligence and machine learning are increasingly integrated into DevOps to automate tasks and enhance predictive analytics.",
		"The adoption of DevSecOps emphasizes incorporating security measures throughout the DevOps pipeline."}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Want %v Got %v", want, err)
	}
}

func TestOutputFile(t *testing.T) {

	tests := []struct {
		searchString      string
		fileName          string
		outputFile        string
		outputFileContent []string
		expectedOutput    string
	}{
		{"DevOps", "test_1.txt", "out_2.txt", []string{"Artificial intelligence and machine learning are increasingly integrated into DevOps to automate tasks and enhance predictive analytics.",
			"The adoption of DevSecOps emphasizes incorporating security measures throughout the DevOps pipeline."}, ""},
		{"DevOps", "test_1.txt", "out.txt", []string{}, "file 'out.txt' already exists"},
	}

	for _, test := range tests {

		file, _ := os.Open(test.fileName)

		defer file.Close()

		output := writeToFile(file, test.searchString, test.outputFile)
		if output != nil && !strings.Contains(output.Error(), test.expectedOutput) {
			t.Errorf("Want %v Got %v", test.expectedOutput, output)
		}

	}

}
