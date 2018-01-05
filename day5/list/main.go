package main

import "fmt"

type People struct{
    Age int
    Name string
}


type Student struct{
    Age int
    Name string 
    Next *People
}


func testList(){
    var s Student
    s.Age = 10
    s.Name = "abc"
    // s.Next = &People{
    //     Age:12, 
    //     Name:"def",
	// }
	s.Next = new(People)
	s.Next.Name = "zj"
	fmt.Printf("s:%+v\n",s)
	fmt.Printf("s.Next:%v\n",*(s.Next))
}

func main(){
	testList()
}