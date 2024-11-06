// Sinta Sarwo - 2311102132

package main

import (
    "fmt"
)

const NMAX int = 127

type tabel struct {
    tab [NMAX]rune
    m   int
}

func isiArray(t *tabel, n *int) {
    var x string
    fmt.Print("Teks: ")
    fmt.Scanln(&x)

    *n = len(x)
    for i := 0; i < *n && i < NMAX; i++ {
        t.tab[i] = rune(x[i])
    }
}

func cetakArray(t tabel, n int) {
    for i := 0; i < n; i++ {
        fmt.Print(string(t.tab[i]))
    }
    fmt.Println()
}

func balikanArray(t *tabel, n int) {
    for i := 0; i < n/2; i++ {
        t.tab[i], t.tab[n-i-1] = t.tab[n-i-1], t.tab[i]
    }
}

func isPalindrome(t tabel, n int) bool {
    for i := 0; i < n/2; i++ {
        if t.tab[i] != t.tab[n-i-1] {
            return false
        }
    }
    return true
}

func main() {
    var tab tabel
    var m int

    isiArray(&tab, &m)

    fmt.Print("Reverse teks: ")
    balikanArray(&tab, m)
    cetakArray(tab, m)

    hasil := isPalindrome(tab, m)

    fmt.Print("Palindrom? ", hasil)
}
