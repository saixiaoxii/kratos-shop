GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	cd app/cmdb && mkdir -p bin/ && export SW_AGENT_NAME=cmdb && go build -ldflags "-X main.Version=$(shell git describe --tags --always)"  -o ./bin/ ./... && cd ../..
	cd app/aconprom && mkdir -p bin/ && export SW_AGENT_NAME=aconprom && go build -ldflags "-X main.Version=$(shell git describe --tags --always)" -o ./bin/ ./... && cd ../..
	cd app/atreus && mkdir -p bin/ && export SW_AGENT_NAME=atreus && go build -ldflags "-X main.Version=$(shell git describe --tags --always)" -o ./bin/ ./... && cd ../..
	#cd app/message && mkdir -p bin/ && export SW_AGENT_NAME=message && go build -ldflags "-X main.Version=$(shell git describe --tags --always)" -o ./bin/ ./... && cd ../..
	cd app/ecenter && mkdir -p bin/ && export SW_AGENT_NAME=ecenter && go build -ldflags "-X main.Version=$(shell git describe --tags --always)" -o ./bin/ ./... && cd ../..

build-skywalking:
	cd app/cmdb && mkdir -p bin/ && export SW_AGENT_NAME=cmdb && go build -ldflags "-X main.Version=$(shell git describe --tags --always)"  -toolexec="/opt/apache-skywalking-go-0.4.0/bin/skywalking-go-agent--linux-amd64" -o ./bin/ ./... && cd ../..
	cd app/aconprom && mkdir -p bin/ && export SW_AGENT_NAME=aconprom && go build -ldflags "-X main.Version=$(shell git describe --tags --always)" -toolexec="/opt/apache-skywalking-go-0.4.0/bin/skywalking-go-agent--linux-amd64" -o ./bin/ ./... && cd ../..
	cd app/atreus && mkdir -p bin/ && export SW_AGENT_NAME=atreus && go build -ldflags "-X main.Version=$(shell git describe --tags --always)" -toolexec="/opt/apache-skywalking-go-0.4.0/bin/skywalking-go-agent--linux-amd64" -o ./bin/ ./... && cd ../..
	cd app/message && mkdir -p bin/ && export SW_AGENT_NAME=message && go build -ldflags "-X main.Version=$(shell git describe --tags --always)" -toolexec="/opt/apache-skywalking-go-0.4.0/bin/skywalking-go-agent--linux-amd64" -o ./bin/ ./... && cd ../..
	cd app/ecenter && mkdir -p bin/ && export SW_AGENT_NAME=ecenter && go build -ldflags "-X main.Version=$(shell git describe --tags --always)" -toolexec="/opt/apache-skywalking-go-0.4.0/bin/skywalking-go-agent--linux-amd64" -a -o ./bin/ ./... && cd ../..

build-image:
	cd app/cmdb && cp -rf ./../../deploy/build/Dockerfile-build Dockerfile &&docker build . -t 100.64.10.37/devops/cmdb:$(shell git describe --tags --always) && cd ../..
	cd app/aconprom && cp -rf ./../../deploy/build/Dockerfile-build Dockerfile &&docker build . -t 100.64.10.37/devops/aconprom:$(shell git describe --tags --always) && cd ../..
	cd app/atreus && cp -rf ./../../deploy/build/Dockerfile-build Dockerfile &&docker build . -t 100.64.10.37/devops/atreus:$(shell git describe --tags --always) && cd ../..
	cd app/message && cp -rf ./../../deploy/build/Dockerfile-build Dockerfile &&docker build . -t 100.64.10.37/devops/message:$(shell git describe --tags --always) && cd ../..
	cd app/ecenter && cp -rf ./../../deploy/build/Dockerfile-build Dockerfile &&docker build . -t 100.64.10.37/devops/ecenter:$(shell git describe --tags --always) && cd ../..

push:
	docker push 100.64.10.37/devops/cmdb:$(shell git describe --tags --always)
	docker push 100.64.10.37/devops/aconprom:$(shell git describe --tags --always)
	docker push 100.64.10.37/devops/atreus:$(shell git describe --tags --always)
	docker push 100.64.10.37/devops/message:$(shell git describe --tags --always)
	docker push 100.64.10.37/devops/ecenter:$(shell git describe --tags --always)
	docker tag 100.64.10.37/devops/cmdb:$(shell git describe --tags --always) 192.168.9.37/devops/cmdb:$(shell git describe --tags --always) 
	docker tag 100.64.10.37/devops/aconprom:$(shell git describe --tags --always) 192.168.9.37/devops/aconprom:$(shell git describe --tags --always)
	docker tag 100.64.10.37/devops/atreus:$(shell git describe --tags --always) 192.168.9.37/devops/atreus:$(shell git describe --tags --always)
	docker tag 100.64.10.37/devops/message:$(shell git describe --tags --always) 192.168.9.37/devops/message:$(shell git describe --tags --always)
	docker tag 100.64.10.37/devops/ecenter:$(shell git describe --tags --always) 192.168.9.37/devops/ecenter:$(shell git describe --tags --always)
	docker push 192.168.9.37/devops/cmdb:$(shell git describe --tags --always) 
	docker push 192.168.9.37/devops/aconprom:$(shell git describe --tags --always)
	docker push 192.168.9.37/devops/atreus:$(shell git describe --tags --always)
	docker push 192.168.9.37/devops/message:$(shell git describe --tags --always)
	docker push 192.168.9.37/devops/ecenter:$(shell git describe --tags --always)

image-del:
	docker rmi 100.64.10.37/devops/cmdb:$(shell git describe --tags --always)
	docker rmi 100.64.10.37/devops/aconprom:$(shell git describe --tags --always)
	docker rmi 100.64.10.37/devops/atreus:$(shell git describe --tags --always)
	docker rmi 100.64.10.37/devops/message:$(shell git describe --tags --always)
	docker rmi 100.64.10.37/devops/ecenter:$(shell git describe --tags --always)

update:
	docker-compose -f deploy/docker/docker-compose-apps.yaml --env-file=.env up -d

update-k8s:
	kubectl set image deployment/cmdb cmdb=192.168.9.37/devops/cmdb:$(shell git describe --tags --always)
	kubectl set image deployment/aconprom aconprom=192.168.9.37/devops/aconprom:$(shell git describe --tags --always)
	kubectl set image deployment/atreus atreus=192.168.9.37/devops/atreus:$(shell git describe --tags --always)
	kubectl set image deployment/message message=192.168.9.37/devops/message:$(shell git describe --tags --always)
	kubectl set image deployment/ecenter ecenter=192.168.9.37/devops/ecenter:$(shell git describe --tags --always)

.PHONY: generate
# generate
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: all
# generate all
all:
	make api;
	make config;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
