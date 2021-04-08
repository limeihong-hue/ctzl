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
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
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

func f() {
	defer fmt.Println("D")
	fmt.Println("F")
}

func Interview48() {
	f()
	fmt.Println("M")
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}
