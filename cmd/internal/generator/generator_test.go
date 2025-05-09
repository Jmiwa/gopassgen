// internal/generator/generator_test.go
package generator

import "testing"

func TestGenerate(t *testing.T) {
    pass, err := Generate(12)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(pass) != 12 {
        t.Errorf("expected password length 12, got %d", len(pass))
    }
}
