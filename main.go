package goisvalidmod

import (
	"fmt"
	"reflect"
)

type Sample struct{
	Name string `required:"true" max:"10"`
	Age int `required:"tidak" max:"2"`
	Address string `required:"true" max:"32"`
}

type IsValidCase struct{
	Name string
	IsValid bool
}

func IsValid(data interface{})[]IsValidCase {
	fields := []IsValidCase{}
	t := reflect.TypeOf(data)
	for i:=0;i < t.NumField();i++{
		field := t.Field(i)
		fmt.Println(field)
		if field.Tag.Get("required") == "true" {
			fields = append(fields, IsValidCase{
				Name:field.Name,
				IsValid: reflect.ValueOf(data).Field(i).Interface() != "",
			})
		
		}else{

			fields = append(fields,IsValidCase{
				Name:field.Name,
				IsValid: true,
			})
		}
	}
	
	return fields
}

func SayHello(name string,age string) string{
	return "hello " + name + age
}