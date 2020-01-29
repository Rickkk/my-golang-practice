package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    bs, err := ioutil.ReadFile("9-func.go")
    if err != nil {
        return
    }
    str := string(bs)
    fmt.Println(str)
}