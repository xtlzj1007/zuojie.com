package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)
	personDB["1"] = PersonInfo{"1", "Tom", "Room 203..."}
	personDB["2"] = PersonInfo{"2", "Jack", "Room 101..."}
	person, ok := personDB["1"]

	//fmt.Println(person.Name,ok)
	if ok {
		fmt.Println("Found person:",person.Name,"With ID 1")
	} else {
		fmt.Println("Did not find person with ID 1")
	}
}
