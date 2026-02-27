package main

import "fmt"

func Convert(s string) int {
    if s == "" {
            return 0
    }

    values := map[byte]int{
            'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100,'D': 500, 'M': 1000,
    }

    total := 0
    n := len(s)

    for i := 0; i < n; i++ {
        current, ok := values[s[i]]

        if !ok {
            fmt.Println("Ошибка")
            return 0
        }

        if i+1 < n {
            next, ok := values[s[i+1]]
            if !ok {
                return 0
            }
            if current < next {
                total -= current
            } else {
               total += current
            }
        } else {
            total += current
        }
    }
    return total
}