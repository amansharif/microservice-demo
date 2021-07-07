package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type data struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	Comment   string `json:"comment"`
}

func getData() (d data) {
	fileBytes, err := ioutil.ReadFile("./data.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &d)

	if err != nil {
		panic(err)
	}

	return d
}

func saveData(d data) () {

	dataBytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./data.json", dataBytes, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Print("Successfully saved")
}
