package functions

import "fmt"

func didStudentPassTheExam(studentName string, finalScore int, totalAttendance int) bool {
	var pass bool

	target := finalScore >= 75 && totalAttendance >= 200

	if target {
		pass = true
	}

	return pass
}

func Returns() {
	name := "Alexa"

	result := didStudentPassTheExam(name, 60, 245)

	if !result {
		fmt.Println("Unfortunately,", name, "didn't pass the exam")
	} else {
		fmt.Println("Congratulations,", name, "passed the exam")
	}
}
