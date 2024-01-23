package udemyLesson

import (
	"fmt"
	"os/user"
	"strconv"
	"strings"
	"time"
)

func buzz() {
	fmt.Println("Buzz !!")
	/*
		multiline comment.
	*/
}
func sector1() {
	// commentout

	/*
		<sample output>
		$go run sample.go
		init !! testtest
		Hello world !!
		Buzz !!
	*/
	fmt.Println("Hello world !!", time.Now())
	fmt.Println(user.Current())
	buzz()
}

func sector2() {
	var var1 int = 1
	var f64 float64 = 1.234
	var s string = "sample"
	var b bool = false

	var (
		var01 int = 1
		var02 int = 2
		var03 int
	)

	xint := 1 // 型宣言なし
	xstring := "sample"

	fmt.Println(var1, f64, s, b)
	fmt.Println(var01, var02, var03)
	fmt.Println(xint, xstring)
	fmt.Printf("%T\n", xstring) // Check variable type
}

func sector3() {
	const Pi = 3.14
	const (
		username = "user"
		password = "pass"
	)

	var (
		u8  uint8   = 255
		i8  int8    = 127
		f32 float32 = 0.123
	)

	fmt.Println(Pi, username, password)
	fmt.Println(u8, i8, f32)
	fmt.Printf("Type is %T Valuse is %v\n", f32, f32)

	fmt.Println("10/3=", 10/3)
	fmt.Println("10.0/3=", 10.0/3)
	fmt.Println("10%3=", 10%3)

	var1 := 0
	var1++
	println(var1)
}

func sector4() {
	var hello string = "Hello"
	var world string = "World"
	helloworld := hello + world

	fmt.Println(helloworld)
	fmt.Println(helloworld[0])
	fmt.Println(string(helloworld[0]))

	// helloworld[0] = "X" <- Bad!!
	fmt.Println(strings.Replace(helloworld, "l", "X", 2))
	fmt.Println(strings.Contains(helloworld, "World"))
	fmt.Println(string(helloworld[0]))

	// how to output ["]
	fmt.Println("\"")
	fmt.Println(`"`)
}

func sector5() {
	t, f := true, false
	fmt.Printf("%T %v %t \n", t, t, t)
	fmt.Printf("%T %v %t \n", f, f, f)
	fmt.Printf("%T %v %t \n", 1, 1, 1)
	fmt.Printf("%T %v %t \n", 0, 0, 0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func sector6() {
	var x int = 123
	xx := float64(x)
	fmt.Printf("%T %v %t \n", xx, xx, xx)

	// convert string to integer
	var s string = "14"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)
}

func sector7() {
	var a [2]int
	a[0] = 100
	a[1] = 200

	var b [2]int = [2]int{100, 200}
	// append(b, 300) -> Occur error!!
	fmt.Println(b)

	var c []int = []int{100, 200}
	c = append(c, 300)
	fmt.Println(c)

}

func sector8() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n, ",", n[2], ",", n[2:4], ",", n[:2], ",", n[2:])

	n[2] = 100
	fmt.Println(n)

	var board = [][]int{
		[]int{0, 1, 2, 3},
		[]int{4, 5, 6, 7},
		[]int{8, 9, 10, 11},
	}
	fmt.Println(board)

	n = append(n, 100, 200, 300, 400, 500)
	fmt.Println(n)

	// 長さ以外にもキャパの最大値（最大の長さ)も設定可能
	// キャパは初期値から動的に増加する
	n2 := make([]int, 3, 5)
	fmt.Printf("length = %d ,capacity = %d ,value = %v\n", len(n2), cap(n2), n2)
	n2 = append(n2, 0, 0)
	fmt.Printf("length = %d ,capacity = %d ,value = %v\n", len(n2), cap(n2), n2)
	n2 = append(n2, 1)
	fmt.Printf("length = %d ,capacity = %d ,value = %v\n", len(n2), cap(n2), n2)

	n3 := make([]int, 3)
	fmt.Printf("length = %d ,capacity = %d ,value = %v\n", len(n3), cap(n3), n3)
	n3 = append(n3, 0)
	fmt.Printf("length = %d ,capacity = %d ,value = %v\n", len(n3), cap(n3), n3)
}
func sector9() {
	n := make([]int, 5)
	for i := 0; i < 5; i++ {
		n = append(n, i)
		fmt.Println(n)
	}
	fmt.Println(n)

	n2 := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		n2 = append(n2, i)
		fmt.Println(n2)
	}
	fmt.Println(n2)
}

func sector10() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	m["banana"] = 300
	m["newitem"] = 500
	fmt.Println(m)

	// 見つかったかどうかの判定を無視することも可能
	nothing := m["nothingitem"]

	fmt.Println(nothing)
	// m["nothingitem"] = 300
	nothingitem, ok := m["nothingitem"]
	if ok == false {
		fmt.Println("Can not find")
	} else {
		fmt.Println("Find item >>", nothingitem)
	}

	var m2 map[string]int
	if m2 == nil {
		fmt.Println("m2 is nil")
	}
}

func sector11() {
	b := []byte{72, 73}
	fmt.Println(b)
	fmt.Println(string(b))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))
}

func sector12() {
	res1, res2 := calc(12, 13)
	fmt.Println(res1, res2)

	res2 = calcItem(200, 2)
	fmt.Println(res2)

	f := func(i int) {
		fmt.Println("inner function", 1)
	}
	f(1)

	func(i int) {
		fmt.Println("Innner func2")
	}(1)
}

func calc(x int, y int) (int, int) {
	fmt.Println("add function")
	return x + y, x - y
}

func calcItem(price int, num int) (result int) {
	result = price * num
	return
}

func sector13() {
	counter := incrementGenerator()

	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	circle1 := circleAreaGenerator(3.14)
	fmt.Println(circle1(2.0))

	circle2 := circleAreaGenerator(3)
	fmt.Println(circle2(2.0))
}

func incrementGenerator() func() int {
	x := 0
	increment := func() int {
		x++
		return x
	}
	return increment
}

func circleAreaGenerator(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func sector14() {
	foo()
	foo(10)
	foo(10, 20, 30)

	// Array型のsを展開してfoo関数へ渡す
	s := []int{1, 2, 3}
	foo(s...)
}

// 可変長引数
func foo(params ...int) {
	fmt.Println(len(params), params)

	for _, param := range params {
		fmt.Println(param)
	}
}

func Lesson1() {
	// sector14()
}
