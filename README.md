# JHStructMapper Go!

Sample Package code with GoLang for **mapping field from source struct to the target struct**.  There are options to map only some field, all field or custom map too.


# How to Install

> go get https://github.com/kazekim/jhstructmapper-go

## Example

Please import github.com/kazekim/jhstructmapper-go in your go file and then code like this

    type Model struct {  
	   A string  
	  B int  
	  C float64  
	  D *string  
	  S SubModel  
	}
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
    type Test struct {  
	   Param1 string `map:"A"`  
	   B int  
	   C float64  
	   Param4 *string `map:"D"`  
	   S SubModel `map:"S"`  
	}  
  
	var test Test  
  
	err := jhstructmapper.ParseSameFieldName(model, &test)  
	if err != nil {  
	   panic(err)  
	}  
	fmt.Println(test)


There is option to map field only field with tag "map". It means that field with no tag will not be mapped.

	type Test struct {  
	   Param1 string `map:"A"`  
	   Param2 int `map:"B"`  
	   Param3 float64 `map:"C"`  
	   Param4 *string
	   S SubModel  
	}
	var test2 Test
	err = jhstructmapper.ParseWithMapTag(model, &test2)
	fmt.Println(test)
