run-app:
	go run cmd/app/main.go

run-worker:
	go run cmd/worker/main.go

run-dispatcher:
	go run cmd/dispatcher/main.go

build:
	docker build -t clodoaldomarques/balances-api:$(version) -f scripts/docker/app/Dockerfile .
	docker build -t clodoaldomarques/balances-api:latest -f scripts/docker/app/Dockerfile .

push:
	docker push clodoaldomarques/balances-api:latest
	docker push clodoaldomarques/balances-api:$(version)

publish:
	docker build -t clodoaldomarques/balances-api:$(version) -f scripts/docker/app/Dockerfile .
	docker build -t clodoaldomarques/balances-api:latest -f scripts/docker/app/Dockerfile .
	docker push clodoaldomarques/balances-api:latest
	docker push clodoaldomarques/balances-api:$(version)

test:
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out