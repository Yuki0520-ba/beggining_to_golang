package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func sector2_1() {
	num := 4
	if num%2 == 0 {
		fmt.Println("ok")
	} else if num%3 == 0 {
		fmt.Println("by3")
	} else {
		fmt.Println("bad")
	}

	x, y := 10, 20
	if x == 10 && y == 20 {
		fmt.Println("ok and")
	}
	if x == 20 || y == 20 {
		fmt.Println("ok or")
	}

	if result := by2(10); result == "ok" {
		fmt.Println("by2")
	} else {
		fmt.Println("not by2")
	}
}

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "bad"
	}
}

func sector2_2() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		} else if i >= 7 {
			break
		}
		fmt.Println(i)
	}

	sum := 1
	for sum < 10 {
		fmt.Println(sum)
		sum += sum
	}
	for {
		fmt.Println("無限ループ")
	}
}

func sector2_3() {
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println(k)
	}
	for _, v := range m {
		fmt.Println(v)
	}

}

func sector2_4() {
	os := "mac"
	switch os {
	case "mac":
		fmt.Println("mac")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println("linux")
	}

	switch os = getOSname(); os {
	case "mac":
		fmt.Println("mac")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println("linux")
	}

	t := time.Now()
	fmt.Println(t.Hour())

	switch {
	case t.Hour() < 12:
		fmt.Println("moning")
	case t.Hour() < 17:
		fmt.Println("aftetrnoon")
	case t.Hour() < 24:
		fmt.Println("night")
	}
}

func getOSname() string {
	return "linux"
}

func sector2_5() {
	defer fmt.Println("world")
	defer fmt.Println("!!")
	fmt.Println("hello")

	file, _ := os.Open("./README.md")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))

}

func sector2_6() {

	loggSettings("logfile.log")
	_, err := os.Open("./sanpleFile")
	if err != nil {
		log.Fatalln("Can not open.")
	}

	log.Println("logging")
	log.Printf("%T %v", "test", "test")
	log.Fatalln("Exit!!")
	log.Println("after")
}

func loggSettings(logfile string) {
	file, _ := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	mylmultiLogFile := io.MultiWriter(os.Stdout, file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(mylmultiLogFile)
}

func sector2_7() {
	file, err := os.Open("README.md")
	if err != nil {
		log.Fatal("Error")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal("Read error")
	}
	fmt.Println(count, string(data))
}

func lesson2() {
	sector2_7()
}
