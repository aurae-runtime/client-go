default: all

all: generate build

build:
	@go build ./...

generate: generate-go generate-go-grpc

generate-go:
	@for API_VERSION in $$(ls ./aurae/api | grep -v ".md"); do \
  		for PROTO_FILE in $$(ls "./aurae/api/$${API_VERSION}" | grep ".proto"); do \
  			protoc \
            		-I"./aurae/api/$${API_VERSION}" \
            		--go_out=./pkg/api/$${API_VERSION}/$${PROTO_FILE%.*} \
            		--go_opt=paths=source_relative \
            		./aurae/api/$${API_VERSION}/$${PROTO_FILE}; \
  		done; \
  	done

generate-go-grpc:
	@for API_VERSION in $$(ls ./aurae/api | grep -v ".md"); do \
  		for PROTO_FILE in $$(ls "./aurae/api/$${API_VERSION}" | grep ".proto"); do \
  			protoc \
            		-I"./aurae/api/$${API_VERSION}" \
            		--go-grpc_out=./pkg/api/$${API_VERSION}/$${PROTO_FILE%.*} \
            		--go-grpc_opt=paths=source_relative \
            		./aurae/api/$${API_VERSION}/$${PROTO_FILE}; \
  		done; \
  	done

submodules: submodule ## Alias for submodule
submodule: ## Initialize all submodules
	@echo "Initializing submodules"
	@echo ""
	@read -p "Warning: This will destroy all work in subdirectories! Press any key to continue." FOO

	# aurae
	@if [ -d /tmp/aurae ]; then rm -rvf /tmp/aurae; fi
	@if [ -d aurae ]; then mv -v aurae /tmp/aurae; fi

	# Init and update
	@git submodule update --init --recursive
	@git submodule update --remote --rebase

	# Attach to main
	cd aurae && git checkout main && git branch && git pull origin main