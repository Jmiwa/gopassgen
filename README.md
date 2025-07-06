# gopassgen

![License](https://img.shields.io/badge/license-MIT-blue.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/Jmiwa/gopassgen)](https://goreportcard.com/report/github.com/Jmiwa/gopassgen)
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.15324155.svg)](https://doi.org/10.5281/zenodo.15324155)
[![Coverage Status](https://coveralls.io/repos/github/Jmiwa/gopassgen/badge.svg?branch=main)](https://coveralls.io/github/Jmiwa/gopassgen?branch=main)
[![build](https://github.com/Jmiwa/gopassgen/actions/workflows/build.yaml/badge.svg)](https://github.com/Jmiwa/gopassgen/actions/workflows/build.yaml)





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
