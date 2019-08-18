package main

import (
	"fmt"
	"github.com/kazekim/jhstructmapper-go"
)

type Model struct {
	A string
	B int
	C *float64
	D *string
	S SubModel
}

type SubModel struct {
	E string
}

func main() {

	s := "This is a pointer"
	f := 12.22
	model := Model{
		"Kim",
		18,
		&f,
		&s,
		SubModel{
			"Test",
		},
	}

	// This will success
	type Test struct {
		Param1 string `map:"A"`
		B int
		C *float64
		Param4 *string `map:"D"`
		S SubModel `map:"S"`
	}

	var test Test

	err := jhstructmapper.ParseSameFieldName(model, &test)
	if err != nil {
		panic(err)
	}
	fmt.Println(test)

	var test3 Test
	err = jhstructmapper.ParseWithMapTag(model, &test3)
	if err != nil {
		panic(err)
	}
	fmt.Println(test3)


	//This is a fail example
	type TestFail struct {
		Param1 string `map:"A"`
		Param2 int `map:"B"`
		Param3 *float64 `map:"C"`
		Param4 *string `map:"D"`
		S SubModel `map:"G"`
	}
	//Assign S to map with G but not field G in Model struct

	var test2 TestFail
	err = jhstructmapper.ParseWithMapTag(model, &test2)

	// It will show error here
	if err != nil {
		panic(err)
	}
	fmt.Println(test)

}