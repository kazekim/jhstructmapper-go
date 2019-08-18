package main

import (
	"fmt"
	"github.com/kazekim/jhstructmapper-go"
)

type Test struct {
	Param1 string `map:"A"`
	Param2 int `map:"B"`
	Param3 float64 `map:"C"`
	Param4 *string `map:"D"`
	S SubModel `map:"S"`
}

type Model struct {
	A string
	B int
	C float64
	D *string
	S SubModel
}

type SubModel struct {
	E string
}

func main() {

	s := "This is a pointer"
	model := Model{
		"Kim",
		18,
		12.22,
		&s,
		SubModel{
			"Test",
		},
	}

	// This will success
	var test Test

	err := jhstructmapper.ParseStruct(model, &test)
	if err != nil {
		panic(err)
	}
	fmt.Println(test)


	//This is a fail example
	type TestFail struct {
		Param1 string `map:"A"`
		Param2 int `map:"B"`
		Param3 float64 `map:"C"`
		Param4 *string `map:"D"`
		S SubModel `map:"G"`
	}
	//Assign S to map with G but not field G in Model struct

	var test2 TestFail
	err = jhstructmapper.ParseStruct(model, &test2)

	// It will show error here
	if err != nil {
		panic(err)
	}
	fmt.Println(test)

}