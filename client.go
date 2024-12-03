package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Student struct {
	Age    int     `json:"age"`
	Name   string  `json:"name"`
	Grades []Grade `json:grades`
}

type Grade struct {
	History int `json:"history"`
	Math    int `json:"math"`
	Science int `json:"science"`
}

func GetStudentsWithStruct(c *http.Client) ([]Student, error) {
	url := "localhost:8080/students"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	var s []Student
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	json.Marshal()

}

func computeAverageMath() float32 {
	return 0.0
}
