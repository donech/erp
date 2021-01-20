PROJECT:=erp
CONFIG:=app.yaml

.PHONY: run
run:
	@echo "wellcome for donech land, this will run the erp system for you."
	@go run main.go -c $(CONFIG)

.PHONY: build
build:
	go build -o bin/$(PROJECT) main.go