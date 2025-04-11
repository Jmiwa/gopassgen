# gopassgen

**gopassgen** is a simple CLI tool written in Go to generate secure passwords with custom options.

## 🔧 Features

- Set password length
- Include digits, symbols, uppercase and lowercase characters
- Generate multiple passwords at once
- Easy to use from your terminal

## 🚀 Usage

```bash
# Generate a 16-character password with digits and symbols
gopassgen --length 16 --digits --symbols

# Generate 5 random passwords, 20 characters long
gopassgen --length 20 --count 5
