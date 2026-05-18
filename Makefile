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


kube-secrets:
	kubectl create secret generic mysql-balances --from-literal=root='a1s2d3f4' --from-literal=balances='b4l4nc3s'
	kubectl create secret generic aws-balances --from-literal=AWS_ACCESS_KEY_ID='test' --from-literal=AWS_SECRET_ACCESS_KEY='test' --from-literal=AWS_SESSION_TOKEN='' --from-literal=aws-account='000000000000' --from-literal=aws-assume-role='' --from-literal=aws-region='us-east-1'

kube-create:
	kubectl apply -f scripts/k8s/mysql-service.yaml
	kubectl apply -f scripts/k8s/localstack-service.yaml
	kubectl apply -f scripts/k8s/redis-service.yaml
	kubectl apply -f scripts/k8s/app-service.yaml

kube-delete:
	kubectl delete -f scripts/k8s/mysql-service.yaml
	kubectl delete -f scripts/k8s/localstack-service.yaml
	kubectl delete -f scripts/k8s/redis-service.yaml
	kubectl delete -f scripts/k8s/app-service.yaml

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