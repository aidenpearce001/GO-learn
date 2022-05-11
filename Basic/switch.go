package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	switch {
		case 0 < now.Hour() && now.Hour() < 12:
			fmt.Println("Its morning")
		case now.Hour() == 12:
			fmt.Println("Its noon")
		case 12 < now.Hour() && now.Hour() < 18:
			fmt.Println("Its afternoon")
		default:
			fmt.Println("Its Night")
	}

	WhatType := func(i interface{}) {
		switch t:= i.(type){
		case bool:
			fmt.Println("Im bool")
		case int:
			fmt.Println("Im int")
		default:
			fmt.Println("Dont know type", t)
		}
	}

	WhatType(true)
    WhatType(1)
    WhatType("hey")
}