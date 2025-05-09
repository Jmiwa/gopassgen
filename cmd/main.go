package main

import (
    "fmt"
    "gopassgen/internal/generator"
)

func main() {
    password, err := generator.Generate(12)
    if err != nil {
        panic(err)
    }
    fmt.Println(password)
}

