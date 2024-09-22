.PHONY: run-local
run-local:
	@echo "Running the application locally..."
	go run main.go ./env/dev/.env.auth

.PHONY: setup-local-db
setup-local-db:
	@echo "Starting local database..."
	docker compose -f docker-compose.db.yml up --build --force-recreate -d
