package main

import (
	"fmt"
	"reflect"
)

func main(){
	slc := []bool{true, true, false, true, false, false}

	literal_slices := [] struct {
		x int
		b string
	}{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}

	fmt.Println(slc)
	fmt.Println(literal_slices)

	// v := reflect.ValueOf(literal_slices)

	// values := make([]interface{}, v.Field())

    // for i := 0; i < v.NumField(); i++ {
    //     values[i] = v.Field(i).Interface()
    // }

	// fmt.Println(values)

}