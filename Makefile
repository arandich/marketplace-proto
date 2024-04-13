default: help

.PHONY: help
help:
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

.PHONY: install-go-deps
install-go-deps: # Install GoLang dependencies for protoc
	@echo "[ - ] Installing golang dependencies"
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@PATH="${PATH}:${HOME}/go/bin"

.PHONY: compile-go
compile-go: install-go-deps  # Compile only protobuf files for GoLang
	@buf mod update
	@buf generate

.PHONY: compile-proto
compile-proto: # Compile all proto files. Only GoLang supported
	@echo "[ - ] Please pre-install nodejs using nvm and protoc"
	@make compile-go
