package main

import (
    "os"
)

func main() {
    file, err := os.Create("test.txt")
    if err != nil {
        // здесь перехватывается ошибка
        return
    }
    defer file.Close()

    file.WriteString("test")
}