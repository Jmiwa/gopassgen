// internal/generator/generator.go
package generator

import (
    "crypto/rand"
    "math/big"
)

const (
    lower       = "abcdefghijklmnopqrstuvwxyz"
    upper       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    digits      = "0123456789"
    symbols     = "!@#$%^&*()-_=+[]{}<>?/|"
)

func Generate(length int, useDigits bool, useSymbols bool) (string, error) {
    charset := lower + upper
    if useDigits {
        charset += digits
    }
    if useSymbols {
        charset += symbols
    }

    if len(charset) == 0 {
        return "", nil
    }

    password := make([]byte, length)
    for i := range password {
        n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
        if err != nil {
            return "", err
        }
        password[i] = charset[n.Int64()]
    }
    return string(password), nil
}
