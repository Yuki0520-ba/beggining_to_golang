package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"

	mygreet "mypkg/greeting" // <モジュール名>/<パッケージへの相対パス> で記載。

	greeting "github.com/tenntenn/greeting/v2"
)

func Go_exercise_01(method string) {
	// 1から100までで奇数のものと偶数ものを出力する
	for i := 0; i < 100; i++ {
		if method == "for" {
			if i%2 == 0 {
				println(i, "-奇数")
			} else {
				println(i, "-偶数")
			}
		} else if method == "switch" {
			a := i % 2
			switch a {
			case 1:
				fmt.Println(i, "-奇数")
			case 0:
				fmt.Println(i, "-偶数")
			}
		}
	}
}

func Go_exercise_02() {
	// １から6の値をランダムに出して、値によって大吉から凶のどれかを出力
	t := time.Now().UnixNano()

	// Go1.20から以下の処理は非推奨
	// rand.Seed(t)

	rand.NewSource(t)
	rand_int := rand.Intn(6) + 1

	switch rand_int {
	case 6:
		fmt.Println("大吉")
	case 5, 4:
		fmt.Println("中吉")
	case 3, 2:
		fmt.Println("吉")
	case 1:
		fmt.Println("凶")
	}
}

func swap(x int, y int) (int, int) {
	return y, x
}
func swap_with_pointer(xp *int, yp *int) {
	tmp := *xp
	*xp = *yp
	*yp = tmp
}
func Go_exercise_03() {
	// Swapの実装
	n, m := swap(10, 20)
	fmt.Println(n, m)

	// ポインタを利用して変数の値を入れ替える
	x, y := 10, 20
	fmt.Println("Befor >>", x, y)
	swap_with_pointer(&x, &y)
	fmt.Println("After >>", x, y)
}

func Go_exercise_04() {
	msg := mygreet.Do() // インポートした内部パッケージ内の関数を実行
	fmt.Println(msg)
	fmt.Println(mygreet.Test_var)

	// 内部パッケージのインポート時にハマった
	// 以下記事を参考に内部モジュールのインポート方法を確認して解決した。
	// https://qiita.com/fetaro/items/31b02b940ce9ec579baf
}

func Go_exercise_05() {
	// インストールした外部パッケージ内の関数を実行
	msg := greeting.Do(time.Now())
	fmt.Println(msg)
}

type Stringer interface {
	String() string
}

func ToStringer(v interface{}) (Stringer, error) {
	s, is_stringer := v.(Stringer) // 変数vがStringerインタフェースを満たしているかどうか
	if is_stringer {
		return s, nil
	} else {
		return nil, MyErr("Not Stringter")
	}
}

type str string // string型に関数をセットすることができないため、独自の文字列形式の方を定義する。

func (s str) String() string {
	return string(s)
}

type MyErr string

func (e MyErr) Error() string {
	return fmt.Sprintf("Error >> %s", e)
}

func Go_exercise_06() {
	// Stringerインターフェースに変換する関数を実装して、
	// エラーが起きないかをチェックする
	my_str := str("hoge hoge")
	msg, err := ToStringer(my_str)
	if err != nil {
		fmt.Println("error ouured at contert to string.>> ", err)
	} else {
		fmt.Println("Converted msg is", msg)
	}

	// 以下の処理は変数my_stringがStringerインタフェースを満たしていないため、
	// ToStringer関数でエラーとなってしまう
	my_string := "not contvert variable."
	msg, err = ToStringer(my_string)
	if err != nil {
		fmt.Println("error ouured at contert to string.>> ", err)
	} else {
		fmt.Println("Converted msg is", msg)
	}
}

type Scanner struct {
	io  io.Reader
	buf [16]byte
}

func (s *Scanner) Scan() (rune, error) {
	n, err := s.io.Read(s.buf[:])
	if err != nil {
		return 0, err
	}

	r, size := utf8.DecodeRune(s.buf[:n])
	if r == utf8.RuneError {
		return 0, errors.New("RuneError ocuured")
	}
	s.io = io.MultiReader(bytes.NewReader(s.buf[size:n]))
	return r, nil
}

func CreateScanner(r io.Reader) *Scanner {
	return &Scanner{io: r}
}

func Go_exercise_07() {
	// 1コードポイント(rune)ずつ読み込むScannerを作る。
	// 初期化時にio.Readerを渡す
	// 備忘：理解ができていていない状態で実装してしまっている。io.Readerの仕組みは。runeとは。低レイヤーの理解が必要か。
	scnnr := CreateScanner(strings.NewReader("Hello, 世界"))
	for {
		r, err := scnnr.Scan()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%c\n", r)
	}

	//  複数のエラーを結合しておくことができる。
	err1 := errors.New("Error1")
	err2 := errors.New("Error2")
	err3 := errors.Join(err1, err2)
	fmt.Errorf("error is %w", err3)

}

func practice_for_array_and_slice() {
	var ns1 = [5]int{1, 2, 3, 4, 5} // 配列に該当。要素数を明示的に宣言してメモリを確保
	var ns2 = []int{1, 2, 3, 4, 5}  // スライスに該当。要素数を明示的に宣言しない。動的に割り当てる
	for i := range ns1 {
		fmt.Printf("Nmber is %d \n", i)
	}
	for i := range ns2 {
		fmt.Printf("Nmber is %d \n", i)
	}

	ns3 := ns1[2:4]
	// ns3 = ns3[0:100]                  // Punicが起きてしまう(cap以上の要素数を確保できない) -> runtime error: slice bounds out of range [:100] with capacity 3
	fmt.Println(ns3)                     // [3 4]
	fmt.Printf("Max is %d \n", len(ns3)) // Max is 2
	fmt.Printf("Max is %d \n", cap(ns3)) // Max is 3  <-切り出し元の配列（またはスライス）の容量が用いられる

	var var2 struct {
		name string
		num  int
	}
	var2.name = "name"
	var2.num = 1
	fmt.Println(var2.name, var2.num)
}

func practice_for_map() {
	var1 := map[string]int{}
	var1["a"] = 1
	var1["b"] = 2
	var1["c"] = 3
	for i := 0; i < 3; i++ {
		n, is_found := var1["a"]
		if is_found {
			fmt.Printf("key a is %d \n", n)
			delete(var1, "a")
		} else {
			fmt.Println("not found.")
		}
	}

	var2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for key, value := range var2 {
		fmt.Printf("key is %s, value is %d. \n", key, value)
	}
}

func practice_for_my_type() {
	// My_structというStructを拡張した型を自身で定義
	type My_struct struct {
		name  string
		value string
	}

	var1 := My_struct{name: "john", value: "one"}
	var2 := My_struct{name: "Jimmy", value: "two"}

	fmt.Println(var1.name, var1.value)
	fmt.Println(var2.name, var2.value)
}

func practice_for_func() {
	msg := "Hello world"
	// こちらのクロージャーが関数内で実行される
	func(m string) {
		fmt.Printf("Msg is %s \n", m)
	}(msg)

	// 関数型の変数を定義して実行する
	vars := []string{"hoge", "fuga"}
	function_vars := []func() string{}
	for _, v := range vars {
		// Goのクロージャは、変数が参照されるときにその変数の最新の値を持つ。
		// そのため、ループ変数 v が最終的に "fuga" となり、全ての関数が同じ値を返すことになることから、
		// ループのスコープ内で定義したローカル変数を利用する。
		local_v := v
		f := func() string { return local_v }
		function_vars = append(function_vars, f)
	}
	for _, f := range function_vars {
		fmt.Println("Function result >> ", f())
	}
}

type Person struct {
	name string
	age  int
}

func (p Person) say_my_name() string {
	return p.name
}
func (p *Person) add_my_age(year int) {
	p.age += year
}
func practice_for_method_and_reciever() {
	// Person型の変数を作成してオブジェクト指向でのクラスメソッドに該当する「メソッド」say_my_nameを実行
	john := Person{
		name: "John",
		age:  10,
	}
	fmt.Println("My name is", john.say_my_name())

	// Person型の変数内の要素を直接書き換える
	// オブジェクト指向でいうところのクラス変数を関数を通して書き換えるイメージ
	fmt.Println("Befor my age is", john.age) // 10
	john.add_my_age(3)
	fmt.Println("After my age is", john.age) // 13
}

func run_panic(msg string) {
	panic("Panic at sub routine >> " + msg)
}
func practice_for_panic_and_recover() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()
	run_panic("test error.")
	panic("Panic at main routine.") // こちらの処理はなされない。run_paic関数ないで最初にパニックが起きてしまい、その時点で終了される

}

func Exercise() {
	// Go_exercise_01("switch")
	// Go_exercise_02()
	// Go_exercise_03()
	// Go_exercise_04()
	// Go_exercise_05()
	// Go_exercise_06()
	// Go_exercise_07()
	// Go_exercise_08() // at go_routine_excecise.go
	Go_exercise_09() // at go_routine_excecise.go

	// practice_for_array_and_slice()
	// practice_for_map()
	// practice_for_my_type()
	// practice_for_func()
	// practice_for_method_and_reciever()
	// practice_for_interface()
	// practice_for_embedded_struct()
	// practice_for_panic_and_recover()
	// practice_for_chanel()
	// practice_for_lock_and_unlock()

}
