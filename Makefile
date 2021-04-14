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

foundEnvoy := $(shell docker ps -f "name=envoy" -q | grep -q . && echo Found || echo Not Found)

.PHONY: envoy
ifeq ($(foundEnvoy), Found)
envoy: stop-envoy start-envoy
	echo "stop and start envoy success"
else
envoy: start-envoy
	echo "start envoy success"
endif

.PHONY: stop-envoy
stop-envoy:
	docker stop envoy

.PHONY: start-envoy
start-envoy:
	docker run -itd --rm --name envoy \
    	  -p 51051:51051 \
    	  -p 10000:10000 \
          -v "$(shell pwd)/internal/proto/service_definition.pb:/data/service_definition.pb:ro" \
          -v "$(shell pwd)/envoy/envoy-config-dev.yml:/etc/envoy/envoy.yaml:ro" \
          envoyproxy/envoy

.PHONY: in-envory
in-envoy:
	docker exec -it envoy /bin/bash

.PHONY: valid-envoy
valid-envoy:
	docker run --rm \
              -v "$(shell pwd)/internal/proto/service_definition.pb:/data/service_definition.pb:ro" \
              -v "$(shell pwd)/envoy/envoy-config-dev.yml:/etc/envoy/envoy.yaml:ro" \
              envoyproxy/envoy \
              --mode validate -c /etc/envoy/envoy.yaml
