package main

import(

 "fmt"
)
type person struct {
	name string
}

func main() {
	var m map[person]int
	p := person{"make"}
	fmt.Println(m[p])
}

func main() {
	 a := [5]int{1, 2, 3, 4, 5 }
	 t := a[3:4::4]
	 fmt.Println(t[0])
}