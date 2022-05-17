package main

import "fmt"

type coord struct {
	lat, long float64
}

var m map[string]coord

func main() {
	dict1:=make(map[string]int)

	dict1["x"] = 12
	dict1["y"] = 7

	fmt.Println(dict1)

	x := dict1["x"]
	fmt.Println(x)

	delete(dict1, "y")
	fmt.Println(dict1)

	key,value := dict1["x"]
	fmt.Println(key)
	fmt.Println(value)

	dict2:=map[string]int{"a":12,"b":23}
	fmt.Println(dict2)

	m = make(map[string]coord)
	m["London"] = coord{40.68433,-65.5345435,}

	fmt.Println(m)
	fmt.Println(m["London"])
}