
.ONESHELL:
SHELL := /bin/sh
set-minikube-docker-daemon:
	$(eval "minikube docker-env")

reset-docker-daemon:
	$(eval "minikube docker-env -u")

build-client:
	docker build -f build/client/dockerfile -t ori-service/location-checkin-client .

build-server:
	docker build -f build/server/dockerfile -t ori-service/location-checkin-server .

deploy-server:
	kubectl apply -f deployments/k8s/server/

deploy-client:
	kubectl apply -f deployments/k8s/client/