App     := gopassgen
Version := `grep '^const VERSION = ' cmd/version.go | sed "s/^const VERSION = \"\(.*\)\"/\1/"`

# ◆◆◆ CI 用レシピ ◆◆◆

update_version new_version = "":
	if [ "{{new_version}}" != "" ]; then \
		sed 's/\$VERSION/{{new_version}}/g' .template/README.md    > README.md; \
		sed 's/\$VERSION/{{new_version}}/g' .template/version.go  > cmd/version.go; \
	fi

make_distribution_files:
	for os in linux windows darwin; do \
		for arch in amd64 arm64; do \
			dir="dist/$(App)-$$os-$$arch"; \
			mkdir -p $$dir; \
			env GOOS=$$os GOARCH=$$arch \
				go build -o $$dir/$(App) cmd/main.go; \
			cp README.md LICENSE $$dir; \
			tar czf $$dir.tar.gz -C dist $(basename $$dir); \
		done; \
	done

upload_assets tag:
	gh release upload \
		--repo Jmiwa/$(App) \
		$tag dist/*.tar.gz




