package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)


//Structs from JSON
type Expression struct {
	Operation string `json:"Operation"`
	Units     Units  `json:"Units"`
}
type Units struct {
	One int `json:"one"`
	Two int `json:"two"`
}


func main() {
	//Reading in the JSON File
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	//Retreiving Data from JSON and putting into struct
	var theData []Expression
	err2 := json.Unmarshal([]byte(data), &theData)
	if err2 != nil{
		fmt.Println(err)
	}

	//Performing the operations
	fileAppend, err := os.OpenFile("data.json", os.O_APPEND|os.O_WRONLY, 0644)
	for _,i := range theData {
		theOperation := i.Operation
		switch theOperation{
		case "Add":
			sum := i.Units.One + i.Units.Two
			fileAppend.WriteString("Addition: " + strconv.Itoa(sum) + "\n")
			break
		case "Subtract":
			difference := i.Units.One - i.Units.Two
			fileAppend.WriteString("Subtraction: " + strconv.Itoa(difference) + "\n")
			break
		case "Multiply":
			product := i.Units.One * i.Units.Two
			fileAppend.WriteString("Multplication: " + strconv.Itoa(product) + "\n")
			break
		case "Divide":
			quotient := i.Units.One / i.Units.Two
			fileAppend.WriteString("Division: " + strconv.Itoa(quotient) + "\n")
			break
		default:
			fmt.Println("Invalid Operation")
		}
	}
	fileAppend.Close()
}