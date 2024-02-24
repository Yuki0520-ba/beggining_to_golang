package simpleTestPractice

import (
	"errors"
	"fmt"
)

func simpleFizzBuzz(inputNum int) (string, error) {
	if inputNum < 0 {
		return "", errors.New("input number is must be higher than 0. ")
	}
	if inputNum%3 == 0 && inputNum%5 == 0 {
		return "FizzBuzz", nil
	}
	if inputNum%3 == 0 {
		return "Fizz", nil
	}

	if inputNum%5 == 0 {
		return "Buzz", nil
	}

	return "", nil
}

func FizzBuzz() {
	for i := 1; i < 101; i++ {
		res, err := simpleFizzBuzz(i)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Count %v is %v \n", i, res)
	}
}
