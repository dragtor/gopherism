package main

import "testing"

type TestCase struct {
    Input string
    ExpectedOutput int
}

var TestCases = []TestCase{
    TestCase{
      Input: "oneTwoThree",
      ExpectedOutput : 3,
    },
}

func TestCamelCase(t *testing.T){
   for _, tc := range TestCases {
    n := countCamelCase(tc.Input)
    if tc.ExpectedOutput != n {
         t.Errorf("input : %s , ExpectedOutput : %d, actual output : %d \n",tc.Input,tc.ExpectedOutput,n)
    }
}
}
