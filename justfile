@default: help
App := 'gopassgen'
Version := `grep '^const VERSION = ' cmd/main/version.go | sed "s/^VERSION = ¥"¥(.*¥)¥"/¥1/g"`
# show help message
@help:
    echo "Build tool for {{ App }} {{ Version }} with Just"
    echo "Usage: just <recipe>"
    echo ""
    just --list

# ビルドとテストの実行
build: test
    go build -o gopassgen /main/main.go

# テストの実行
test:
go test -covermode=count -coverprofile=coverage.out ./...