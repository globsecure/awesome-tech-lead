.PHONY: generate-readme setup

setup:
	@echo "Instalando dependências..."
	go mod download
	@echo "Dependências instaladas com sucesso!"

generate-readme:
	@go run cmd/generate_readme/main.go

lint:
	@go run cmd/lint/lint.go

all: setup generate-readme
