# ------------------ Golang commands ------------------
.PHONY: clean install update test build run
clean:
	@go clean -cache
	@go clean -testcache
	@go clean -modcache
install:
	@go mod download
update:
	@go get -u all
test:
	@go test -v ./cmd/test/...
build:
	@go build -a -installsuffix cgo -o ./.build/main ./cmd/main.go
run:
	@./.build/main

# Third party libraries commands
.PHONY: scan
scan:
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@govulncheck ./...

# ------------------ Docker commands ------------------
.PHONY: docker-redis docker-postgres docker-api
docker-redis:
	@docker compose --env-file app.env up --build --no-deps -d redis
docker-postgres:
	@docker compose --env-file app.env up --build --no-deps -d postgres
docker-api:
	@docker compose --env-file app.env up --build --no-deps -d api

.PHONY: docker-ghcr-login
docker-ghcr-login:
	@echo "" ;\
	echo "Login - GitHub Docker Registry" ;\
	read -p "Enter your Github username: " gUsername ;\
	read -p "Enter your Github personal access token: " gPass ;\
	echo "" ;\
	echo $$gPass | docker login ghcr.io -u $$gUsername --password-stdin ;\

.PHONY: docker-ghcr-push-specific docker-ghcr-pull-specific
docker-ghcr-push-specific:
	@echo "" ;\
	echo "Tag - GitHub Docker Registry" ;\
	gCorp="emenec-finance" ;\
	gRepo="school-api" ;\
	read -p "Enter your package name(redis, postgres, api): " gPackage; gTag=$${gPackage:-"api"} ;\
	read -p "Enter your tag(default is 1): " gTag; gTag=$${gTag:-"1"} ;\
	docker tag go-api-$$gPackage ghcr.io/$$gCorp/$$gRepo/$$gPackage:$$gTag ;\
	echo "" ;\
	echo "Pushing $$gPackage - GitHub Docker Registry" ;\
	docker push ghcr.io/$$gCorp/$$gRepo/$$gPackage:$$gTag;
docker-ghcr-pull-specific:
	@echo "" ;\
	echo "Tag - GitHub Docker Registry" ;\
	gCorp="emenec-finance" ;\
	gRepo="school-api" ;\
	read -p "Enter your package name(redis, postgres, api): " gPackage; gTag=$${gPackage:-"api"} ;\
	read -p "Enter your tag(default is 1): " gTag; gTag=$${gTag:-"1"} ;\
	echo "" ;\
	echo "Pulling $$gPackage - GitHub Docker Registry" ;\
	docker pull ghcr.io/$$gCorp/$$gRepo/$$gPackage:$$gTag
