# gopassgen

**gopassgen** は、Goで実装されたシンプルなコマンドライン用パスワード生成ツールです。  
指定したオプションに応じて、安全でカスタマイズ可能なパスワードを生成できます。

## 🔧 特徴

- パスワードの長さを指定可能
- 数字・記号・大文字・小文字を含めるか選択可能
- 複数のパスワードを一度に生成
- CLIツールとして簡単に使用可能

## 🚀 使い方（Usage）

```bash
# 長さ16文字、数字と記号を含むパスワードを生成
gopassgen --length 16 --digits --symbols

# 長さ20文字のパスワードを5つ生成
gopassgen --length 20 --count 5


## 🚀 Usage

```bash
# Generate a 16-character password with digits and symbols
gopassgen --length 16 --digits --symbols

# Generate 5 random passwords, 20 characters long
gopassgen --length 20 --count 5
