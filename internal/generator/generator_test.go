// internal/generator/generator_test.go
package generator

import (
    "testing"
)

func TestGenerateDefaultCharset(t *testing.T) {
    password, err := Generate(12, false, false) // 英字のみ
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(password) != 12 {
        t.Errorf("expected password length 12, got %d", len(password))
    }
}

func TestGenerateWithDigits(t *testing.T) {
    password, err := Generate(16, true, false)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(password) != 16 {
        t.Errorf("expected password length 16, got %d", len(password))
    }
}

func TestGenerateWithSymbols(t *testing.T) {
    password, err := Generate(20, false, true)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(password) != 20 {
        t.Errorf("expected password length 20, got %d", len(password))
    }
}

func TestGenerateWithDigitsAndSymbols(t *testing.T) {
    password, err := Generate(24, true, true)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(password) != 24 {
        t.Errorf("expected password length 24, got %d", len(password))
    }
}
