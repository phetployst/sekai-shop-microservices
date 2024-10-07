.PHONY: run-local run-auth run-player run-inventory run-item run-payment

RUN_CMD = @echo "Running the application with env file: $1"; go run main.go $1

run-local: run-auth run-player run-inventory run-item run-payment

run-auth:
	$(call RUN_CMD, ./env/dev/.env.auth)

run-player:
	$(call RUN_CMD, ./env/dev/.env.player)

run-inventory:
	$(call RUN_CMD, ./env/dev/.env.inventory)

run-item:
	$(call RUN_CMD, ./env/dev/.env.item)

run-payment:
	$(call RUN_CMD, ./env/dev/.env.payment)

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
	