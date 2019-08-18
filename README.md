# JHStructMapper Go!

Sample Package code with GoLang for **mapping field from source struct to the target struct**.  There are options to map only some fields or map all fields.


# How to Install

> go get https://github.com/kazekim/jhstructmapper-go

## Example

Please import github.com/kazekim/jhstructmapper-go in your go file and then code like this

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
    // Output : {Kim 18 0xc0000160f8 0xc0000101f0 {Test}}


There is option to map field only field with tag "map". It means that field with no tag will not be mapped.

    err = jhstructmapper.ParseWithMapTag(model, &test2)
    if err != nil {  
      panic(err)  
    } 
    fmt.Println(test)
 
    //Output {Kim 0 <nil> 0xc000010200 {Test}}
