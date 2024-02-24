package simpleTestPractice

import "testing"

func TestSimpleFizzBuzz(t *testing.T) {
	testData := []struct {
		name  string
		val   int
		want  string
		isErr bool
	}{
		{name: "test for 0", val: 0, want: "FizzBuzz", isErr: false},
		{name: "test for not  division with 1", val: 1, want: "", isErr: false},
		{name: "test for Fizz with 3", val: 3, want: "Fizz", isErr: false},
		{name: "test for Buzz with 5", val: 5, want: "Buzz", isErr: false},
		{name: "test for fizzBuzz with 15", val: 0, want: "FizzBuzz", isErr: false},
		{name: "test for error with minus", val: -3, want: "", isErr: true},
	}
	for _, test := range testData {
		test := test
		t.Run(test.name, func(t *testing.T) {
			res, err := simpleFizzBuzz(test.val)
			isError := (err != nil)
			if res != test.want && isError != test.isErr {
				t.Errorf("input value : %v, want value: %v,reslut value: %v, want error: %v, result error: %v", test.val, test.want, res, test.isErr, isError)
			} else {
				t.Logf("test case %v is passed.", test.name)
			}
		})
	}
}

// Fuzzing Test
func FuzzTestForSimpleFizzBuzz(f *testing.F) {
	f.Fuzz(func(f *testing.T, i int) {
		_, _ = simpleFizzBuzz(i)
	})
}
