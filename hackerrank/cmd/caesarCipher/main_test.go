package main

import "testing"

type TestCase struct {
    S string 
    K int32
    ExpectedOutput string
}

var TestCases = []TestCase{
    TestCase{
         S: "abcdefghijklmnopqrstuvwxyz",
         K: int32(2),
         ExpectedOutput: "cdefghijklmnopqrstuvwxyzab",
    },
}

func TestCaesarCipher(t *testing.T) {
    for _ , tc := range TestCases {
        actualOutput := caesarCipher(tc.S,tc.K)
        if tc.ExpectedOutput != actualOutput {
            t.Errorf("input : %s ,Expected Output : %s , Actual output : %s\n ",tc.S,tc.ExpectedOutput, actualOutput )
        }
         


}}

