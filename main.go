package main

import (
	"fmt"
	"net/http"
)

func main() {
	go runServer()
	c := http.DefaultClient
	fmt.Println("Printing students with struct:")
	students, err := GetStudentsWithStruct(c)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf("%#v\n\n", students)
	fmt.Println("Printing students with map string interface:")
	studentsInterface, err := GetStudentsWithInterface(c)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf("%v\n", studentsInterface)

	avg := calcAvgGradeStruct(*students)
	avgInterface := calcAvgGradeInterfacePrecise(studentsInterface)
	fmt.Printf("%v\n", avg)
	fmt.Printf("%v\n", avgInterface)
}
