package main

import (
	"fmt"
	"os"
	"strconv"
	"io/ioutil"
	"encoding/json"
	. "session3/models"
)

func printData(no int) {
	file, _ := ioutil.ReadFile("data.json")

	var persons []Person
	json.Unmarshal([]byte(file), &persons)

	if no > len(persons) {
		fmt.Println("data not found")
	} else {
		index := no-1
		fmt.Println("Nama:", persons[index].Name)
		fmt.Println("Alamat:", persons[index].Address)
		fmt.Println("Pekerjaan:", persons[index].Job)
		fmt.Println("Alasan memilih kelas Golang:", persons[index].Reason)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("required 1 argument: <integer>")
		return
	}

	input := os.Args[1]
	no, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("argument must be integer")
		return
	}

	printData(no)
}
