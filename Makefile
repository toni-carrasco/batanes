dev:
	@LP_MODE=serve LP_PORT=3001 air

serve:
	@LP_MODE=serve go run cmd/main.go

build:
	@go run cmd/main.go

install-devtools:
	@echo "Installing recommended devtools"
	@go install github.com/air-verse/air@latest
	@echo "Devtools installed"
