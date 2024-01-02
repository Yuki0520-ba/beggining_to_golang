package main

import (
	"fmt"
	"math/rand"
	"time"

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

// ////////////////
type nameSpeakerInterface interface {
	say_my_name() string
}
type Japanese struct {
	name           string
	favorite_comic string
}

func (j Japanese) say_my_name() string {
	return "私の名前は" + j.name + "です。"
}

type American struct {
	name                string
	favorite_nba_player string
}

func (a American) say_my_name() string {
	return "My name is " + a.name + "."
}

func SelfIntroduction(human nameSpeakerInterface) {
	fmt.Println(human.say_my_name())
}

func practice_for_interface() {
	taro := Japanese{
		name:           "太郎",
		favorite_comic: "ワンピース",
	}
	john := American{
		name:                "John",
		favorite_nba_player: "Lebron James",
	}

	SelfIntroduction(taro) // result -> 私の名前は太郎です。
	SelfIntroduction(john) // result -> My name is John.
}

type Tokyo_poeple struct {
	//　自身で定義したタイプを組み込むことができる。
	// 組み込んだタイプに定義されたプロパティも利用可能
	address string
	people  Japanese
}

func (t Tokyo_poeple) say_my_address() string {
	return fmt.Sprintf("私は%sに住んでいます。", t.address)
}
func (t Tokyo_poeple) say_my_name() string {
	return t.people.say_my_name()
}

type addressSpeakerInterfce interface {
	say_my_address() string
}

type selfIntroductionInteraface interface {
	// 別のインターフェースを組み込むことができる。
	// 親子関係のような形でインタフェースを定義することができる。
	// 複数のインタフェースが組み込まれている場合、ANDで機能する。
	nameSpeakerInterface
	addressSpeakerInterfce
}

func SelfIntroductionDetails(self selfIntroductionInteraface) {
	fmt.Println(self.say_my_name())
	fmt.Println(self.say_my_address())
}

func practice_for_embedded_struct() {
	hanako := Tokyo_poeple{
		address: "東京都港区",
		people: Japanese{
			name:           "花子",
			favorite_comic: "花男",
		},
	}
	SelfIntroductionDetails(hanako) // result -> 私の名前は花子です。私は東京都港区に住んでいます。

	/*
		以下の処理は「selfIntroductionInteraface」で定義された関数（say_my_addressとsay_my_name）が
		実装されていないTypeを利用しているのでエラーとなる
	*/
	// jimmy := American{
	// 	name:                "Jimmy",
	// 	favorite_nba_player: "none",
	// }
	// SelfIntroductionDetails(jimmy)
}

//////////////////

func Exercise() {
	// Go_exercise_01("switch")
	// Go_exercise_02()
	// Go_exercise_03()
	// Go_exercise_04()
	// Go_exercise_05()

	// practice_for_array_and_slice()
	// practice_for_map()
	// practice_for_my_type()
	// practice_for_func()
	// practice_for_method_and_reciever()
	// practice_for_interface()
	practice_for_embedded_struct()

}
