app-run:
	go run cmd/app/main.go

worker-run:
	go run cmd/worker/main.go

dispatcher-run:
	go run cmd/dispatcher/main.go

app-build:
	docker build -t clodoaldomarques/balances-api:$(version) -f scripts/docker/app/Dockerfile .
	docker build -t clodoaldomarques/balances-api:latest -f scripts/docker/app/Dockerfile .

app-push:
	docker push clodoaldomarques/balances-api:latest
	docker push clodoaldomarques/balances-api:$(version)

publish: app-build app-push

kube-secrets:
	kubectl create secret generic mysql-secrets --from-literal=root='a1s2d3f4' --from-literal=balances='b4l4nc3s'
	kubectl create secret generic aws-secrets --from-literal=AWS_ACCESS_KEY_ID='test' --from-literal=AWS_SECRET_ACCESS_KEY='test' --from-literal=AWS_SESSION_TOKEN='' --from-literal=aws-account='000000000000' --from-literal=aws-assume-role='' --from-literal=aws-region='us-east-1'

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
	terraform -chdir=scripts/terraform/ plan
	terraform -chdir=scripts/terraform/ apply

terraform-init:
	terraform -chdir=scripts/terraform/ init

terraform-destroy:
	terraform -chdir=scripts/terraform/ destroy

minikube: kube-secrets kube-create terraform

test:
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out