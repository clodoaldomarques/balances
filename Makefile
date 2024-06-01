run: kubectl-delete kubectl-apply run-app run-worker run-dispatcher

run-app:
	go run cmd/app/main.go

run-worker:
	go run cmd/worker/main.go

run-dispatcher:
	go run cmd/dispatcher/main.go

build-app:
	docker build -t clodoaldomarques/balances-api:$(version) -f scripts/docker/app/Dockerfile .
	docker build -t clodoaldomarques/balances-api:latest -f scripts/docker/app/Dockerfile .

push-app:
	docker push clodoaldomarques/balances-api:latest
	docker push clodoaldomarques/balances-api:$(version)

publish-app: build-app push-app

publish-all: build-app push-app

mysql-secrets:
	kubectl create secret generic root-pass --from-literal=password='a1s2d3f4'
	kubectl create secret generic balances-pass --from-literal=password='b4l4nc3s'

kubectl-apply:
	kubectl apply -f scripts/k8s/mysql-service.yaml
	kubectl apply -f scripts/k8s/localstack-service.yaml
	kubectl apply -f scripts/k8s/redis-service.yaml

kubectl-delete:
	kubectl delete -f scripts/k8s/mysql-service.yaml
	kubectl delete -f scripts/k8s/localstack-service.yaml
	kubectl delete -f scripts/k8s/redis-service.yaml

terraform-init:
	terraform -chdir=scripts/terraform/ init
terraform-plan:
	terraform -chdir=scripts/terraform/ plan
terraform-apply:
	terraform -chdir=scripts/terraform/ apply

test:
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out