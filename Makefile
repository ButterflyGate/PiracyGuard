.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## linux amd64 用のバイナリ作成
	./shells/build.sh

linux: ## linux amd64 用のバイナリ作成
	./shells/build.sh linux amd

windows: ## windows intel 用のバイナリ作成
	./shells/build.sh windows amd

mac: ## mac intel 用のバイナリ作成
	./shells/build.sh mac amd

linux-arm: ## linux arm64 用のバイナリ作成
	./shells/build.sh linux arm

windows-arm: ## windows arm64 用のバイナリ作成
	./shells/build.sh windows arm

mac-arm: ## mac arm64 用のバイナリ作成
	./shells/build.sh mac arm
