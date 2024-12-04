package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type StudentsResponse struct {
	Students []Student `json:"students"`
}

type Student struct {
	Age    int    `json:"age"`
	Name   string `json:"name"`
	Grades Grades `json:"grades"`
}

type Grades struct {
	History int `json:"history"`
	Math    int `json:"math"`
	Science int `json:"science"`
}

func GetStudentsWithStruct(c *http.Client) (*StudentsResponse, error) {
	url := "http://127.0.0.1:8080/students"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	var s StudentsResponse
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resBody, &s)
	return &s, nil
}

func GetStudentsWithInterface(c *http.Client) (map[string]interface{}, error) {
	url := "http://127.0.0.1:8080/students"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	var s map[string]interface{}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resBody, &s)
	return s, nil
}

func calcAvgGradeStruct(sr StudentsResponse) float32 {
	total := 0.0
	for _, s := range sr.Students {
		total += float64(s.Grades.Math)
	}
	return float32(total / float64(len(sr.Students)))
}

// Panic: interface conversion: interface {} is []interface {}, not []map[string]interface {}
// func calcAvgGradeInterface(st map[string]interface{}) float32 {
// 	total := 0
// 	for _, s := range st["students"].([]map[string]interface{}) {
// 		total += s["grade"].(map[string]int)["math"]
// 	}
// 	return float32(total / len(st["students"].([]map[string]interface{})))
// }

func calcAvgGradeInterfacePrecise(st map[string]interface{}) float32 {
	// If we don't type assert here, we get an error on line 89: cannot range over students (variable of type interface{})
	students, ok := st["students"]
	if !ok {
		fmt.Println("students key not found in map")
		return -1.0
	}
	// Errors in type assertion must be handled seperately to those generated when extracting a value. If these errors are unhandled, the program panics.
	// Next, JSON arrays are unmarshalled into []interface{} first. Each individual element must then be asserted to map[string]interface{}
	studentsInterface, ok := students.([]interface{})
	if !ok {
		fmt.Println("could not convert students to an array of maps")
		return -1.0
	}
	total := 0.0
	for _, st := range studentsInterface {
		s := st.(map[string]interface{})
		if !ok {
			fmt.Println("could not assert student interface to map string intreface")
			return -1.0
		}
		grades, ok := s["grades"]
		if !ok {
			fmt.Printf("student %s does not have a grade", s["name"])
			return -1.0
		}
		fmt.Printf("grades: %v, type: %T\n", grades, grades) // Debugging line
		// Cannot be map[string]int -> numeric values are unmarshalled into float64 by default
		gradesMap, ok := grades.(map[string]interface{})
		if !ok {
			fmt.Printf("student %s grades could not be converted to a map", s["name"])
			return -1.0
		}
		mathGrade, ok := gradesMap["math"].(float64)
		if !ok {
			fmt.Println("could not convert math grade to float64")
			return -1.0
		}
		total += mathGrade

	}
	return float32(total / float64(len(studentsInterface)))
}
