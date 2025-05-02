# Justfile

# バイナリ名とバージョン文字列を定義
App     := gopassgen
Version := `grep '^const VERSION = ' cmd/main/version.go | sed "s/^const VERSION = \"\(.*\)\"/\1/"`

# デフォルト・ヘルプ
@default: help

@help:
	echo "Build tool for {{App}} {{Version}}"
	echo "Usage: just <recipe>"
	echo ""
	just --list

# (1) ビルド（テスト実行込み）
build: test
	go build -o {{App}} cmd/main.go

# (2) テスト + カバレッジレポート
test:
	go test -covermode=count -coverprofile=coverage.out ./...

# (3) クリーンアップ
clean:
	go clean
	rm -f coverage.out {{App}}
	rm -rf dist

# ◆◆◆ CI 用レシピ（資料スライド14 より）◆◆◆

# (4) バージョン更新（new_version が空でなければ置換）
update_version new_version = "":
	if [ "{{new_version}}" != "" ]; then \
		sed 's/\$VERSION/{{new_version}}/g' .template/README.md      > README.md; \
		sed 's/\$VERSION/{{new_version}}/g' .template/version.go > cmd/main/version.go; \
	fi

# (5) 各プラットフォーム向けバイナリ＆アーカイブ作成
make_distribution_files:
	for os in linux windows darwin; do \
		for arch in amd64 arm64; do \
			dir="dist/{{App}}-$$os-$$arch"; \
			mkdir -p $$dir; \
			env GOOS=$$os GOARCH=$$arch \
				go build -o $$dir/{{App}} cmd/main.go; \
			cp README.md LICENSE $$dir; \
			tar czf $$dir.tar.gz -C dist $(basename $$dir); \
		done; \
	done

# (6) GitHub Release へのアセットアップロード
upload_assets tag:
	gh release upload \
		--repo Jmiwa/{{App}} \
		$tag dist/*.tar.gz




