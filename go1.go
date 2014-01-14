package main

import (
    . "fmt"     // Using '.' to avoid prefixing functions with their package names
    . "runtime" //   This is probably not a good idea for large projects...
    . "time"
)

var i = 0

func adder() {
    for x := 0; x < 1000000; x++ {
        i++
    }
}

func decr() {
    for x :=0 ; x<1000000; x++ {
        i--
    }
}

func main() {
    GOMAXPROCS(NumCPU())
    go adder()
    go decr()
    Sleep(100*Millisecond)
    Println("Done:", i);
}
