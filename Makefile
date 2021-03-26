PROJECT:=erp
CONFIG:=app.yaml
APP:=grpc

.PHONY: run
run: gen envoy
	@echo "wellcome for donech land, this will run the erp system for you."
	@go run main.go $(APP) -c $(CONFIG)

.PHONY: build
build: gen
	go build -o bin/$(PROJECT) main.go

.PHONY: gen
gen:
	protoc --go_out=plugins=grpc:./internal/proto/ ./internal/proto/*.proto -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I./internal/proto/
	protoc --proto_path=./internal/proto/ --proto_path=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I./internal/proto/  --include_imports --include_source_info --descriptor_set_out=./internal/proto/service_definition.pb ./internal/proto/*.proto

.PHONY: envoy
envoy:
	docker stop envoy
	docker run -itd --rm --name envoy \
	  -p 51051:51051 \
      -v "$(shell pwd)/internal/proto/service_definition.pb:/data/service_definition.pb:ro" \
      -v "$(shell pwd)/envoy/envoy-config-dev.yml:/etc/envoy/envoy.yaml:ro" \
      envoyproxy/envoy
