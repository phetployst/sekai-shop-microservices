.PHONY: run-local
run-local:
	@echo "Running the application locally..."
	go run main.go ./env/dev/.env.auth
	