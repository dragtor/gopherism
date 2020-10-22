package main

import ("unicode"
"fmt"
)

func caesarCipher(s string, k int32) string {
    var output []rune
    for _ , r := range s {
        var rs rune
        if !unicode.IsLetter(r) {
           output = append(output, r)
           continue
        }
        initial := 'a'
        end := 'z'
        if unicode.IsUpper(r){
            initial = 'A'
            end = 'Z'
        }
            rs = (r + (k %26))
            if rs > end {
                rs = initial +  (rs - end - 1)
            }

            output = append(output,rs)
    }
            fmt.Println(output)
    return string(output)
}

