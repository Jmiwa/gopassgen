# gopassgen

![License](https://img.shields.io/badge/license-MIT-blue.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/<your-username>/gopassgen)](https://goreportcard.com/report/github.com/<your-username>/gopassgen)
https://img.shields.io/badge/:badgeContent?style=flat&logo=MIT&logoColor=blue&label=license


**gopassgen** は、Goで作られたシンプルかつ安全なパスワード生成ツールです。

## 📦 概要

長さや構成（数字・記号・大文字など）を細かく指定して、複数のパスワードを一括生成できます。CLI上で簡単に実行でき、スクリプトやパイプ処理にも組み込みやすい設計を意識しています。

## 🧪 使用例（Usage）

```bash
# デフォルト（長さ12のパスワードを1つ）
gopassgen

# 長さ16、数字と記号を含む
gopassgen --length 16 --digits --symbols

# 5つのパスワードを一括生成
gopassgen --length 20 --count 5
