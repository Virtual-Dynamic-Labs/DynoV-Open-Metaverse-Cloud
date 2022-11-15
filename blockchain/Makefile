.PHONY: run
run: ## run the server
	go run cmd/server/main.go

.PHONY: dep
dep: ## update all bazel file with necessary dependency
	go mod tidy

.PHONE: doc
doc: ## update swagger document
	swag init --parseDependency --parseInternal  -g cmd/server/main.go