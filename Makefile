.PHONY: run-local
run-local:
	@echo "Running the application locally..."
	go run main.go ./env/dev/.env.auth

.PHONY: setup-local-db
setup-local-db:
	@echo "Starting local database..."
	docker compose -f docker-compose.db.yml up --build --force-recreate -d

.PHONY: gen-gRPC-all gen-gRPC-authPb gen-gRPC-playerPb gen-gRPC-itemPb gen-gRPC-inventoryPb

gen-gRPC-all: gen-gRPC-authPb gen-gRPC-playerPb gen-gRPC-itemPb gen-gRPC-inventoryPb

gen-gRPC-%: 
	@echo "Regenerating gRPC code from $*Pb.proto..."
	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	modules/$*/$*Pb/$*Pb.proto
	