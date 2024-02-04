package main

import (
	"mypkg/goBasicalPractice"
	"mypkg/simpleHttpServer"
	"mypkg/udemyLesson"
)

func init() {}

func main() {
	udemyLesson.Lesson1()
	udemyLesson.Lesson2()
	goBasicalPractice.Exercise()
	simpleHttpServer.RunServer()
}
