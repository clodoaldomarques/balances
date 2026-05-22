api = balances-api
worker = balances-worker
repository = clodoaldomarques

run-app:
	go run cmd/main.go

build-app:
	docker build -t $(repository)/$(api):$(version) -f scripts/docker/app/Dockerfile .
	docker tag $(repository)/$(api):$(version) $(repository)/$(api):latest

push-app:
	docker push $(repository)/$(api):$(version)
	docker push $(repository)/$(api):latest

publish-app: build-app push-app	

worker-run:
	go run cmd/worker/main.go

dispatcher-run:
	go run cmd/dispatcher/main.go

kube-apply:
	kubectl apply -f scripts/k8s/

kube-destroy:
	kubectl delete -f scripts/k8s/ --ignore-not-found

kube-restart: kube-destroy kube-apply

terraform:
	until nc -z 192.168.49.2 30002; do echo waiting for localstack; sleep 2; done;
	terraform -chdir=scripts/terraform/ plan
	terraform -chdir=scripts/terraform/ apply -auto-approve

terraform-init:
	terraform -chdir=scripts/terraform/ init

terraform-destroy:
	terraform -chdir=scripts/terraform/ destroy

test:
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out