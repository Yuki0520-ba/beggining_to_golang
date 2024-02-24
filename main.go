package main

import (
	"mypkg/goBasicalPractice"
	"mypkg/simpleHttpServer"
	"mypkg/simpleTestPractice"
	"mypkg/udemyLesson"
)

func init() {}

func main() {
	udemyLesson.Main()
	goBasicalPractice.Exercise()
	simpleHttpServer.RunServer()
	simpleTestPractice.Main()
}
