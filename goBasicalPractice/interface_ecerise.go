package goBasicalPractice

import "fmt"

type nameSpeakerInterface interface {
	say_my_name() string
}
type addressSpeakerInterfce interface {
	say_my_address() string
}

type selfIntroductionInteraface interface {
	// 別のインターフェースを組み込むことで
	// 親子関係のような形でインタフェースを定義することができる。
	// 複数のインタフェースが組み込まれている場合、ANDで機能する。
	nameSpeakerInterface
	addressSpeakerInterfce
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

func SelfIntroduction(human nameSpeakerInterface) {
	fmt.Println(human.say_my_name())
}
func SelfIntroductionDetails(self selfIntroductionInteraface) {
	fmt.Println(self.say_my_name())
	fmt.Println(self.say_my_address())
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
