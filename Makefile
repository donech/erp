PROJECT:=erp
CONFIG:=app.yaml
APP:=grpc
ENVOY:=envoy-donech
JAEGER:=jaeger-donech

.PHONY: run
run: gen
	@echo "wellcome for donech land, this will run the erp system for you."
	@go run main.go $(APP) -c $(CONFIG)

.PHONY: depend
depend: envoy jaeger
	@echo "start depends success"
.PHONY: build
build: gen
	go build -o bin/$(PROJECT) main.go

.PHONY: gen
gen:
	protoc --go_out=plugins=grpc:./internal/proto/ -I./internal/proto/third_party/googleapis ./internal/proto/*.proto -I./internal/proto/
	protoc --proto_path=./internal/proto/ -I./internal/proto/ -I./internal/proto/third_party/googleapis  --include_imports --include_source_info --descriptor_set_out=./internal/proto/service_definition.pb ./internal/proto/*.proto

foundEnvoy := $(shell docker ps -f "name=$(ENVOY)" -q | grep -q . && echo Found || echo Not Found)

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
	docker stop $(ENVOY)

.PHONY: start-envoy
start-envoy:
	docker run -itd --rm --name $(ENVOY) \
    	  -p 51051:51051 \
    	  -p 10000:10000 \
          -v "$(shell pwd)/internal/proto/service_definition.pb:/data/service_definition.pb:ro" \
          -v "$(shell pwd)/envoy/envoy-config-dev.yml:/etc/envoy/envoy.yaml:ro" \
          envoyproxy/envoy

.PHONY: in-envory
in-envoy:
	docker exec -it $(ENVOY) /bin/bash

.PHONY: valid-envoy
valid-envoy:
	docker run --rm \
              -v "$(shell pwd)/internal/proto/service_definition.pb:/data/service_definition.pb:ro" \
              -v "$(shell pwd)/envoy/envoy-config-dev.yml:/etc/envoy/envoy.yaml:ro" \
              envoyproxy/envoy \
              --mode validate -c /etc/envoy/envoy.yaml



foundJaeger := $(shell docker ps -f "name=$(JAEGER)" -q | grep -q . && echo Found || echo Not Found)

.PHONY: jaeger
ifeq ($(foundJaeger), Found)
jaeger: stop-jaeger start-jaeger
	@echo "stop and start jaeger success"
else
jaeger: start-jaeger
	@echo "start jaeger success"
endif
.PHONY: start-jaeger
start-jaeger:
	@echo "jaeger starting"
	docker run --rm -d --name=$(JAEGER) \
				  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
				  -p 9411:9411\
				  -p 16686:16686\
                  jaegertracing/all-in-one
	@echo "jaeger starting"

.PHONY: stop-jaeger
stop-jaeger:
	docker stop $(JAEGER)

