package main

import (
    "fmt"
    "github.com/Jmiwa/gopassgen/internal/generator"
)

func main() {
    password, err := generator.Generate(12)
    if err != nil {
        panic(err)
    }
    fmt.Println(password)
}

