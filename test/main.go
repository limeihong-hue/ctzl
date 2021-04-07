package main

import (
	"fmt"

	"ctzl/test/interview"

	_ "errors"
	"flag"
)

func init() {
	fmt.Println("2")
}

func init() {
	fmt.Println("1")
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

var name string

func main() {
	var n [10]int
	var i, j int

	flag.Parse()

	fmt.Println("name=%s", name)

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
		}

	interview.Interview1()

	Practice1()

	interview.Interview2()

	interview.Interview12()

	interview.Interview14()

	interview.Interview17()

	interview.Interview24()

	interview.Interview25()

	interview.Interview_24()

	interview.Interview26()
	interview.Interview27()

	interview.Interview28()
	interview.Interview30()

	interview.Interview44()
	interview.Interview45()

	interview.Interview47()
	interview.Interview48()2
}
	}

	interview.Interview1()

	Practice1()

	interview.Interview2()

	interview.Interview12()

	interview.Interview14()

	interview.Interview17()

	interview.Interview24()

	interview.Interview25()
}
