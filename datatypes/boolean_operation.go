package datatypes

import "fmt"

func BooleanOperations() {
	var finalScore int8 = 90
	var attendance int8 = 30

	var scorePass bool = finalScore > 75
	var attendancePass bool = attendance > 15

	var isPass bool = scorePass && attendancePass

	fmt.Println(isPass)
}
