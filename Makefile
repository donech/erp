PROJECT:=erp
CONFIG:=app.yaml
APP:=grpc

.PHONY: run
run: gen
	@echo "wellcome for donech land, this will run the erp system for you."
	@go run main.go $(APP) -c $(CONFIG)

.PHONY: build
build: gen
	go build -o bin/$(PROJECT) main.go
