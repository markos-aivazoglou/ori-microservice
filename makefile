
include ./scripts/conf-dev.env
export

unit-test:
	echo 'Running unit tests'
	go test -v ./internal/...

integration-test:
	echo 'Running integration Tests'
run-local-server:
	go run cmd/server/main.go

run-local-client:
	go run cmd/grpc-client/main.go

generate-grpc:
	protoc -I api/proto/ v1/checkin.proto --go_out=plugins=grpc:proto-gen/api_v1

build-all: build-server build-client
	# docker image prune -f

build-client:
	docker build -f build/client/dockerfile -t ori-service/location-checkin-client .

build-server:
	docker build -f build/server/dockerfile -t ori-service/location-checkin-server .

deploy-server:
	kubectl apply -f deployments/k8s/server/

deploy-client:
	kubectl apply -f deployments/k8s/client/
