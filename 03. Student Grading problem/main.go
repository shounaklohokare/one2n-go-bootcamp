package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) []student {

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error occured, ", err)
		os.Exit(1)
	}

	defer file.Close()

	r := csv.NewReader(file)

	headers, err := r.Read()
	if err != nil {
		fmt.Println("Failed to read the headers")
		os.Exit(1)
	}

	fmt.Println("Headers :- ", headers)

	var students []student
	for {

		row, err := r.Read()
		if err != nil {
			break
		}

		var student student

		student.firstName = row[0]
		student.lastName = row[1]
		student.university = row[2]
		student.test1Score, _ = strconv.Atoi(row[3])
		student.test2Score, _ = strconv.Atoi(row[4])
		student.test3Score, _ = strconv.Atoi(row[5])
		student.test4Score, _ = strconv.Atoi(row[6])

		students = append(students, student)

	}

	return students
}

func calculateGrade(students []student) []studentStat {
	var studentStats []studentStat
	for _, student := range students {
		var ss studentStat
		avg := float32(student.test1Score+student.test2Score+student.test3Score+student.test4Score) / 4
		ss.student = student
		ss.finalScore = avg
		ss.grade = getGrade(avg)
		studentStats = append(studentStats, ss)
	}
	return studentStats
}

func getGrade(score float32) Grade {
	if score >= 70 {
		return A
	} else if score < 70 && score >= 50 {
		return B
	} else if score < 50 && score >= 35 {
		return C
	}

	return F
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	topScore := float32(0)
	var topper studentStat
	for _, student := range gradedStudents {
		if student.finalScore > topScore {
			topScore = student.finalScore
			topper = student
		}
	}
	return topper
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	universityToppers := make(map[string]studentStat)
	studentsUniMap := mapStudentsToUniversity(gs)
	for uni, students := range studentsUniMap {
		t := findOverallTopper(students)
		universityToppers[uni] = t
	}
	return universityToppers
}

func mapStudentsToUniversity(gs []studentStat) map[string][]studentStat {
	studentsUniversityMap := make(map[string][]studentStat)
	for _, student := range gs {
		studentsUniversityMap[student.university] = append(studentsUniversityMap[student.university], student)
	}
	return studentsUniversityMap
}
