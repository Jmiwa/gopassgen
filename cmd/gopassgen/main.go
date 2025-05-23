package main

import (
    "flag"
    "fmt"
    "github.com/Jmiwa/gopassgen/internal/generator"
)

func main() {
    // CLIフラグの定義
    length := flag.Int("length", 12, "パスワードの長さ")
    useDigits := flag.Bool("digits", false, "数字を含める")
    useSymbols := flag.Bool("symbols", false, "記号を含める")

    flag.Parse() // フラグを解析

    // パスワード生成
    password, err := generator.Generate(*length, *useDigits, *useSymbols)
    if err != nil {
        panic(err)
    }
    fmt.Println(password)
}
