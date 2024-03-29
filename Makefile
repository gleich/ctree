##########
# Building
##########

build-docker-prod:
	docker build -f docker/Dockerfile -t mattgleich/ctree:latest .
build-docker-dev-lint:
	docker build -f docker/dev.lint.Dockerfile -t mattgleich/ctree:lint .
build-go:
	go get -v -t -d ./...
	go build -v .
	rm ctree

#########
# Linting
#########

lint-golangci:
	golangci-lint run
lint-gomod:
	go mod tidy
	git diff --exit-code go.mod
	git diff --exit-code go.sum
lint-goreleaser:
	goreleaser check
lint-hadolint:
	hadolint docker/Dockerfile
	hadolint docker/dev.lint.Dockerfile
lint-in-docker: build-docker-dev-lint
	docker run mattgleich/ctree:lint

##########
# Grouping
##########

# Linting
local-lint: lint-golangci lint-goreleaser lint-hadolint lint-gomod
docker-lint: lint-in-docker
# Build
local-build: build-docker-prod build-docker-dev-lint
