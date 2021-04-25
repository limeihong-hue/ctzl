package interview

import (
	"fmt"
	"reflect"
	"time"
)

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
		fmt.Println("hello() v", h())
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

func Interview_24() {
	a := [3]int{5, 6, 1}
	b := [3]int{5, 6}

	alen := len(a)
	blen := len(b)

	if alen == blen {
		fmt.Println("len equal")
	} else {
		fmt.Println("len not equal")
	}

	acap := cap(a)
	bcap := cap(b)

	if acap == bcap {
		fmt.Println("cap equal")
	} else {
		fmt.Println("cap not equal")
	}

	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}

type TestStruc struct {
	name string
}

func Interview26() {
	var i interface{} = nil
	if i == nil {
		fmt.Println("nil")
	}
	fmt.Println("not nil")
	var ii interface{} = TestStruc{name: "test"}
	fmt.Printf("%v", ii)
	if ii == nil {
		fmt.Println(" ii nil")
		return
	} else {
		fmt.Println("ii not nil")
	}
}

func Interview27() {
	s := make(map[string]int)
	delete(s, "h")
	fmt.Println("%d,%v", s["h"])

	value, ok := s["h"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("exits map", value)
	} else {
		fmt.Println("do not exist")
	}
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func Interview28() {
	t := Teacher{}
	t.ShowB()

	t.ShowA()
}

func hello30_1(i int) {
	fmt.Println(i)
}

func hello30_2(i *int) {
	fmt.Println(i)
}

func Interview30() {
	i := 5
	defer hello30_1(i)
	defer hello30_2(&i)
	i = i + 10
}

func Interview34() {
	str := "hello"
	//str[0] = 'x'
	str = "abc"
	//str[0] = 'x'

	//str[0] = 'x'

	fmt.Println(str)
}

func increaseA() int {
	var i int
	defer func() {
		i = i + 1
	}()
	return 1
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func Interview44() {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func Interview45() {
	var a A = Work{3}
	s := a.(Work)
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
}

type Person struct {
	age int
}

func Interview47() {
	person := &Person{28}

	// 1.
	defer fmt.Println(person.age)

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29
}

/*
func f1() (r int) {
    defer func() {
        r++
    }()
    return 0
}

func f2() (r int) {
    t := 5
    defer func() {
        t = t + 5
    }()
    return t
}

func f3() (r int) {
    defer func(r int) {
        r = r + 5
    }(r)
    return 1
}
func Interview41() {

	fmt.Println(f1)
}
*/

func Interview41() {
	var m map[string]int = make(map[string]int) //A
	m["a"] = 1
	if v := m["b"]; v != 0 { //B
		fmt.Println(v)
	}
}
func f1() (r int) {

	defer func() {
		r++
	}()
	return 0
}
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
func Interview46() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())

}
func Interview55() {

	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1)
	s2 = append(s2, 5, 6, 7)
	fmt.Println(s1)
}
func Interview56() {
	if a := 1; false {
	} else if b := 2; false {
	} else {
		println(a, b)
	}
}
func Interview58() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

type People1 interface {
	Speak1(string) string
}
type Student struct{}

func (stu Student) Speak1(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func Interview60() {
	var peo People1 = Student{}
	_ = peo
	think := "speak"
	fmt.Println(peo.Speak1(think))
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func Interview63() {
	fmt.Println(South)
}

type People11 interface {
	Show11()
}

type Student11 struct{}

func (stu *Student11) Show11() {

}

type Student21 struct{}

func Interview62() {
	/*
		var s *Student11 //= new(Student11)
		if s == nil {
			fmt.Println("s is nil")
		} else {
			fmt.Println("s is not nil")
		}

		var p People11 = s
		//p := s
		if p == nil {
			fmt.Println("p is nil")
		} else {
			fmt.Println("p is not nil")
		}
	*/

	var s2 *Student21
	_ = s2

	//var p2 interface = s2

	var age int = 5
	_ = age

	var s *Student11
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	var p People11 = s
	//p = 1
	//p = "strintg"

	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}

/*func Interview68() {

    var m = [...]int{1, 2, 3}

    for i, v := range m {
        go func() {
            fmt.Println(i, v)
        }()
    }

    time.Sleep(time.Second * 3)
}
*/

type I1 interface {
	One()
}
type S1 struct {
	I1
	num int
}

func (s1 *S1) One() {
	fmt.Printf("type: S1  value: %+v interface One \n", s1.num)
}

type I2 interface {
	Two()
}
type I12 interface {
	I1
	I2
}
type S12 struct {
	I12
	num int
}

func (s12 *S12) Two() {
	fmt.Printf("type: S12  value: %+v interface Two \n", s12.num)
}

func InterviewXX() {
	var x I1
	s1 := new(S1) //&S1{num: 1}
	s12 := &S12{num: 12}
	_ = s12

	x = s1
	fmt.Printf("1: %T\n", x)             // 1: *main.S1
	fmt.Println("2:", reflect.TypeOf(x)) //2: *main.S1
}

type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": Math{2, 3},
}

func Interview64() {
	tmp := m["foo"]
	tmp.x = 4
	m["foo"] = tmp
	fmt.Println(m["foo"].x)
}

var p1 *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Println(*p1)
}

func Interview66() {
	var err error
	p1, err = foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p1)
}
func Interview67() {
	fmt.Println("Interview67")
	v := []int{1, 2, 3}
	for i, k := range v {
		if i == 0 {
			v[2] = 100
			v[1] = 200
		}
		fmt.Println("i %d …%d", i, k)
		v = append(v, i)
	}
	fmt.Println("Interview67 end")
}

func Interview68() {
	var m = [...]int{1, 2, 3}

	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 3)
}
func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()

	if f == nil {
		dd := 1
		_ = dd
	}
	defer f()
	f = func() {
		r += 2
	}
	//r = n + 1
	return n + 1
}

func Interview69() {
	fmt.Println(f(3))
}
func Interview70() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}
func change(s ...int) {
	s = append(s, 3)
}

func Interview71() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println(slice)
	change(slice[0:2]...)
	fmt.Println(slice)
}
func Interview72() {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

type Foo struct {
	bar string
}

func Interview73() {
	s1 := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}
	s2 := make([]*Foo, len(s1))
	for i, value := range s1 {
		s2[i] = &value
	}
	fmt.Println(s1[0], s1[1], s1[2])
	fmt.Println(s2[0], s2[1], s2[2])
}
func Interview74() {

	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
}
func Interview77() {
	i := 1
	s := []string{"A", "B", "C"}
	i, s[i-1] = 2, "Z"
	fmt.Printf("s: %v \n", s)
}
