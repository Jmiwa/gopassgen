package generator

import (
    "crypto/rand"
    "errors"
    "math/big"
)

const (
    lower   = "abcdefghijklmnopqrstuvwxyz"
    upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    digits  = "0123456789"
    symbols = "!@#$%^&*()-_=+[]{}<>?/|"
)

func randomCharFrom(set string) (byte, error) {
    n, err := rand.Int(rand.Reader, big.NewInt(int64(len(set))))
    if err != nil {
        return 0, err
    }
    return set[n.Int64()], nil
}

func shuffle(bytes []byte) ([]byte, error) {
    for i := len(bytes) - 1; i > 0; i-- {
        jInt, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
        if err != nil {
            return nil, err
        }
        j := int(jInt.Int64())
        bytes[i], bytes[j] = bytes[j], bytes[i]
    }
    return bytes, nil
}

func Generate(length int, useDigits bool, useSymbols bool) (string, error) {
    requiredSets := []string{lower + upper}
    if useDigits {
        requiredSets = append(requiredSets, digits)
    }
    if useSymbols {
        requiredSets = append(requiredSets, symbols)
    }

    if length < len(requiredSets) {
        return "", errors.New("password length too short for required character types")
    }

    charset := lower + upper
    if useDigits {
        charset += digits
    }
    if useSymbols {
        charset += symbols
    }

    password := make([]byte, length)

    // 必須カテゴリから1文字ずつ先に入れる
    for i := 0; i < len(requiredSets); i++ {
        c, err := randomCharFrom(requiredSets[i])
        if err != nil {
            return "", err
        }
        password[i] = c
    }

    // 残りの文字をランダムに埋める
    for i := len(requiredSets); i < length; i++ {
        c, err := randomCharFrom(charset)
        if err != nil {
            return "", err
        }
        password[i] = c
    }

    // シャッフルして順番をランダムに
    shuffled, err := shuffle(password)
    if err != nil {
        return "", err
    }

    return string(shuffled), nil
}
