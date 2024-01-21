package goBasicalPractice

import (
	"bufio"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math/big"
	"os"
	"sync"
	"time"
)

func input(r io.Reader) <-chan string {
	// TODO: チャネルを作る
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			// TODO: チャネルに読み込んだ文字列を送る
			ch <- s.Text()
		}
		// TODO: チャネルを閉じる
		close(ch)
	}()
	// TODO: チャネルを返す
	return ch
}

func Go_exercise_08() {
	// ひたすら標準入力から文字列を受け取り
	// 受け取った文字列を出力する
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		fmt.Println(<-ch)
	}

}

func newRnadomString(num int) string {
	random_int64, _ := rand.Int(rand.Reader, big.NewInt(int64(num)))
	random_n := int(random_int64.Int64())
	runes := make([]byte, random_n)
	for i := 0; i < random_n; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(255))
		runes[i] = byte(num.Int64())
	}
	return base64.RawStdEncoding.EncodeToString(runes)
}

func userTyping(ch chan bool) {
	s := bufio.NewScanner(os.Stdin)
	target_str := newRnadomString(8)
	fmt.Println("■ Target is >>", target_str)
	for s.Scan() {
		ch <- s.Text() == target_str
		if s.Text() == target_str {
			fmt.Println("Corect!!")
		} else {
			fmt.Println("InCorect...")
		}
		target_str = newRnadomString(8)
		fmt.Println("■ Target is >>", target_str)
	}
}

func userTypingGameGoroutine() {
	s := time.Now()
	timer := 30
	ch := make(chan bool)
	go userTyping(ch)

	correct_count := 0
	in_correct_count := 0
	for {
		if <-ch {
			correct_count++
		} else {
			in_correct_count++
		}
		if int((time.Since(s)).Seconds()) >= timer {
			close(ch)
			println("\nFinished!!!")
			break
		}
	}
	fmt.Println("================")
	fmt.Println("Corrct num is >>", correct_count)
	fmt.Println("In Corrct num is >>", in_correct_count)
	fmt.Println("Correct persentage is >>", float32(correct_count)/float32(in_correct_count+correct_count))

}

func userTypingContext(ctx context.Context, is_correct_ch chan bool) {
	s := bufio.NewScanner(os.Stdin)
	target_str := newRnadomString(8)
	fmt.Println("■ Target is >>", target_str)
	for s.Scan() {
		select {
		case <-ctx.Done(): // 制限時間に到達した場合
			close(is_correct_ch)
			return
		default:
			is_correct := s.Text() == target_str
			if is_correct {
				fmt.Println("Corect!!")
			} else {
				fmt.Println("InCorect...")
			}
			is_correct_ch <- is_correct

			target_str = newRnadomString(8)
			fmt.Println("■ Target is >>", target_str)
		}
	}
}

func userTypingGameContext() {
	bc := context.Background()
	timer := 30 * 1 * time.Second
	ctx, cancel := context.WithTimeout(bc, timer)
	defer cancel()

	ch := make(chan bool)
	go userTypingContext(ctx, ch)

	correct_count := 0
	count := 0
	for {
		// ２つ目の返値でチヤネルがオープンになっているかどうかを取得できる
		// 制限時間に到達した場合、呼び出し先のGoroutine側の処理でチャネルがクローズされ
		// 呼び出し元の処理でクローズをキャッチする
		v, is_open := <-ch
		if is_open == false {
			break
		}
		if v {
			correct_count++
		}
		count++
	}
	fmt.Println("================")
	fmt.Println("Corrct num is >>", correct_count)
	fmt.Println("In Corrct num is >>", (count - correct_count))
	fmt.Println("Correct persentage is >>", float32(correct_count)/float32(count))
}

func Go_exercise_09() {
	// タイピングゲーム
	// 標準出力に英単語を出す（出すものは自由）
	// 標準入力から1行受け取る
	// 制限時間内に何問解けたか表示する

	// userTypingGameGoroutine() // goroutineとチャネルのみで実装
	userTypingGameContext() // Contextを利用

}

func practice_for_lock_and_unlock() {
	var l sync.RWMutex
	l.Lock()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("two second after.")
		l.Unlock()
	}()

	l.Lock() // こここれ以上このファンクションの実行が進まないようにロック
	fmt.Println("finished.")
	l.Unlock()

}

func serve(ch chan string, num string) {
	ch <- "first-" + num
	ch <- "second-" + num
	ch <- "third-" + num
	ch <- "finish"
}

func practice_for_chanel() {
	ch := make(chan string)
	ch2 := make(chan string)
	defer close(ch)
	defer close(ch2)

	go serve(ch, "1")
	go serve(ch2, "2")

	break_flag_1 := false
	break_flag_2 := false

	// ↓[sample output]
	// first-2
	// second-2
	// first-1
	// second-1
	// third-1
	// finish
	// third-2
	// finish
	for {
		select {
		case v1 := <-ch:
			fmt.Println(v1)
			if v1 == "finish" {
				break_flag_1 = true
			}
		case v2 := <-ch2:
			fmt.Println(v2)
			if v2 == "finish" {
				break_flag_2 = true
			}
		}
		if break_flag_1 && break_flag_2 {
			break
		}
	}
}
