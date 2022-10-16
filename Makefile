default: all

all: generate build

build:
	@go build ./...

generate: generate-go generate-go-grpc

generate-go:
	@for API_VERSION in $$(ls ./auraed/stdlib | grep ".proto"); do \
  		for PROTO_FILE in $$(ls "./auraed/stdlib/$${API_VERSION}"); do \
  			protoc \
            		-I"./auraed/stdlib/$${API_VERSION}" \
            		--go_out=./pkg/stdlib/$${API_VERSION}/$${PROTO_FILE%.*} \
            		--go_opt=paths=source_relative \
            		./auraed/stdlib/$${API_VERSION}/$${PROTO_FILE}; \
  		done; \
  	done

generate-go-grpc:
	@for API_VERSION in $$(ls ./auraed/stdlib | grep ".proto"); do \
  		for PROTO_FILE in $$(ls "./auraed/stdlib/$${API_VERSION}"); do \
  			protoc \
            		-I"./auraed/stdlib/$${API_VERSION}" \
            		--go-grpc_out=./pkg/stdlib/$${API_VERSION}/$${PROTO_FILE%.*} \
            		--go-grpc_opt=paths=source_relative \
            		./auraed/stdlib/$${API_VERSION}/$${PROTO_FILE}; \
  		done; \
  	done
