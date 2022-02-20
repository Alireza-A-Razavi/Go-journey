package main

import "fmt"

type Student struct {
	name string
	age int
	grade float64
}
func (student Student) IsBetter(s Student) bool {
	return student.grade > s.grade
}



func main(){
	st_1 := Student{"Hamid Reza", 16, 3.2}
	st_2  := Student{
		name: "Mumammad",
		age: 16,
		grade: 3.5,
	}
	statement := st_1.IsBetter(st_2)
	fmt.Println(statement)

	// ----------------------------------------//

	f := FilterFunc(14)
	fmt.Println(f)
}

