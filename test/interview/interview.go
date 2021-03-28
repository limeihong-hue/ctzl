package interview

import "fmt"

func init() {
	fmt.Println("0")
}

func Interview1() {
	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

func Interview2() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

func Interview6() {
	list := new([]int)
	//错误原因：
	//list = append(list, 1)
	fmt.Println(list)
}

func Interview12() {
	a := []int{7, 8, 9}
	fmt.Printf("%+v\n", a)
	ap(a)
	fmt.Printf("%+v\n", a)
	app(a)
	fmt.Printf("%+v\n", a)
}

func ap(a []int) {
	b := append(a, 10)

	b[0] = 1
}

func app(a []int) {
	a[0] = 1
}

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func Interview14() {
	fmt.Println(x, y, z, k, p)
}

func hello() []string {
	return nil
}

func Interview17() {
	h := hello
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
		fmt.Println("hello() %v", h())
	}

}

func hello21(num ...int) {
	num[0] = 18
}

func Interview21() {
	i := []int{5, 6, 7}
	hello21(i...)
	fmt.Println(i[0])
}

func Interview24() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t[0])
	//fmt.Println(t[1])
}

func Interview25() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t[0])
}
