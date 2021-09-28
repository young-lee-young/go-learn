package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id string `json:"id"`
	Name string `json:"NameLee"`
	Class int `json:"class"`
	Address string `json:"address"`
}

func main() {
	student := &Student{
		Id: "447a5f4c-6cef-4ef0-8a0b-3c6556642612",
		Name: "lee",
		Class: 1,
		Address: "beijing",
	}

	jsonString, _ := json.Marshal(student)

	newStudent := new(Student)
	json.Unmarshal(jsonString, newStudent)
	fmt.Println(newStudent)
}
