package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("6-arr.go")
    if err != nil {
        // здесь перехватывается ошибка
        return
    }
    defer file.Close()

    // получить размер файла
    stat, err := file.Stat()
    if err != nil {
        return
    }
    // чтение файла
    bs := make([]byte, stat.Size())
    _, err = file.Read(bs)
    if err != nil {
        return
    }

    str := string(bs)
    fmt.Println(str)
}