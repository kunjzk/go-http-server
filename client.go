package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Students struct {
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

func GetStudentsWithStruct(c *http.Client) (*Students, error) {
	url := "http://127.0.0.1:8080/students"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	var s Students
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

// func computeAverageMath() float32 {
// 	return 0.0
// }
