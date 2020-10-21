package main

import (
	"testing"
)

type TestCase struct {
	Input          string
	ExpectedResult string
	ExpectedError  error
}

var _ = func() bool {
	testing.Init()
	return true
}()

var testcases = []TestCase{
	TestCase{
		Input:          "/",
		ExpectedResult: "https://www.calhoun.io/",
		ExpectedError:  nil,
	},
	TestCase{
		Input:          "https://courses.calhoun.io/signin",
		ExpectedResult: "https://courses.calhoun.io/signin",
		ExpectedError:  nil,
	},
}

func TestIsPathInDomain(t *testing.T) {
	domain := "https://www.calhoun.io/"
	for _, tc := range testcases {
		domain, err := IsPathInDomain(domain, tc.Input)
		if domain != tc.ExpectedResult || err != tc.ExpectedError {
			t.Errorf("input: %s, expectedresult : %s , actualresult: %s, expectederr: %v, actualerr : %v \n", tc.Input, tc.ExpectedResult, domain, tc.ExpectedError, err)
		}

	}
}
