# gopassgen

![License](https://img.shields.io/badge/license-MIT-blue.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/Jmiwa/gopassgen)](https://goreportcard.com/report/github.com/Jmiwa/gopassgen)
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.15324155.svg)](https://doi.org/10.5281/zenodo.15324155)
[![Coverage Status](https://coveralls.io/repos/github/Jmiwa/gopassgen/badge.svg?branch=main)](https://coveralls.io/github/Jmiwa/gopassgen?branch=main)
[![build](https://github.com/Jmiwa/gopassgen/actions/workflows/build.yaml/badge.svg)](https://github.com/Jmiwa/gopassgen/actions/workflows/build.yaml)
![Version](https://img.shields.io/badge/version-1.0.2-blue)




---

**gopassgen** は、Goで作られたシンプルかつ安全なパスワード生成ツールです。

## 📦 概要

長さや構成（数字・記号など）を指定して、複数のパスワードを一括生成できます。CLI上で簡単に実行でき、スクリプトやパイプ処理にも組み込みやすい設計です。

## ✨ 主な特徴

- 任意の長さに対応（デフォルトは12文字）
- 数字・記号の有無を指定可能
- 一括生成機能（複数パスワードを生成）
- シェル補完（bash, zsh, fish, powershell）対応
- Homebrew tap による簡単インストール
- Go製・依存パッケージ最小

## 📥 インストール

### Homebrew（推奨）

```bash
brew install jmiwa/tap/gopassgen
```

### バイナリを直接使用

[Releases](https://github.com/Jmiwa/gopassgen/releases) ページから環境に合ったバイナリをダウンロードしてください。

## 🧪 使用例（Usage）

```bash
# デフォルト（長さ12のパスワードを1つ）
gopassgen

# 長さ16、数字と記号を含む
gopassgen --length 16 --digits --symbols

# 5つのパスワードを一括生成
gopassgen --length 20 --count 5
```

## ⚙️ コマンドオプション一覧

| オプション        | 略記 | 説明                       | デフォルト |
|------------------|------|----------------------------|------------|
| `--length`       | `-l` | パスワードの長さ           | `12`       |
| `--digits`       | `-d` | 数字を含める               | `false`    |
| `--symbols`      | `-s` | 記号を含める               | `false`    |
| `--count`        | `-c` | 生成するパスワードの個数   | `1`        |

## 🧩 シェル補完の設定

各種シェル向けに補完スクリプトを生成可能です。

### zsh

```bash
gopassgen completion zsh > ~/.zsh/completions/_gopassgen
source ~/.zsh/completions/_gopassgen
```

### bash

```bash
gopassgen completion bash > /etc/bash_completion.d/gopassgen
```

### fish

```bash
gopassgen completion fish > ~/.config/fish/completions/gopassgen.fish
```

## 🛠 開発・ビルド方法

### 依存：Just, Go

#### ビルド

```bash
just build
```

#### テスト

```bash
just test
```

## 📁 ディレクトリ構成

```
.
├── cmd               # CLIエントリポイント
├── internal          # ロジック（generatorなど）
├── completions       # 各シェルの補完スクリプト
├── docs              # Hugoによるドキュメント
├── homebrew-tap      # Homebrew formula
└── Justfile          # 開発コマンド管理
```

## 📜 ライセンス

[MIT License](./LICENSE)

---


