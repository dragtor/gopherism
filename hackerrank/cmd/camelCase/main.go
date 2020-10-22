package main

import (
    "strings"
    "unicode"
    //"fmt"
)

func countCamelCase(input string) int{
    f := func(r rune) bool {
        return unicode.IsUpper(r)
    }
    lst := strings.FieldsFunc(input, f)
    return len(lst)
}
