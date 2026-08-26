APP_NAME := ms-url-shortner
BINARY := $(APP_NAME)
DOCKER_IMAGE := $(APP_NAME)
GO := go

HELM_REPO_NAME := ms-url-shortner
HELM_REPO_URL := https://shubham-bansal96.github.io/ms-url-shortner/
HELM_CHART := $(HELM_REPO_NAME)/url-shortner-charts
HELM_RELEASE := url-shortner
NAMESPACE := default

.PHONY: all build run test test-verbose test-coverage clean fmt vet lint vendor tidy docker-build docker-run swagger serve-swagger helm-repo-add helm-install helm-upgrade helm-uninstall helm-status helm-template help

all: fmt vet test build

build:
	CGO_ENABLED=0 $(GO) build -o $(BINARY) .

run:
	$(GO) run main.go

test:
	$(GO) test ./... -count=1

test-verbose:
	$(GO) test ./... -v -count=1

test-coverage:
	$(GO) test ./... -coverprofile=coverage.out -count=1
	$(GO) tool cover -html=coverage.out -o coverage.html

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

lint:
	golangci-lint run ./...

vendor:
	$(GO) mod vendor

tidy:
	$(GO) mod tidy

clean:
	rm -f $(BINARY) main.exe coverage.out coverage.html

docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run: docker-build
	docker run -p 4000:4000 $(DOCKER_IMAGE)

swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

serve-swagger:
	swagger serve -F=swagger swagger.yaml

helm-repo-add:
	helm repo add $(HELM_REPO_NAME) $(HELM_REPO_URL)
	helm repo update

helm-install: helm-repo-add
	helm install $(HELM_RELEASE) $(HELM_CHART) -n $(NAMESPACE)

helm-upgrade: helm-repo-add
	helm upgrade $(HELM_RELEASE) $(HELM_CHART) -n $(NAMESPACE)

helm-uninstall:
	helm uninstall $(HELM_RELEASE) -n $(NAMESPACE)

helm-status:
	helm status $(HELM_RELEASE) -n $(NAMESPACE)

helm-template: helm-repo-add
	helm template $(HELM_RELEASE) $(HELM_CHART) -n $(NAMESPACE)

help:
	@echo "Available targets:"
	@echo "  build           - Build the binary"
	@echo "  run             - Run the application"
	@echo "  test            - Run tests"
	@echo "  test-verbose    - Run tests with verbose output"
	@echo "  test-coverage   - Run tests with coverage report"
	@echo "  fmt             - Format code"
	@echo "  vet             - Run go vet"
	@echo "  lint            - Run golangci-lint"
	@echo "  vendor          - Update vendor directory"
	@echo "  tidy            - Tidy go modules"
	@echo "  clean           - Remove build artifacts"
	@echo "  docker-build    - Build Docker image"
	@echo "  docker-run      - Build and run Docker container"
	@echo "  swagger         - Generate swagger spec"
	@echo "  serve-swagger   - Serve swagger UI"
	@echo ""
	@echo "Kubernetes/Helm:"
	@echo "  helm-repo-add   - Add Helm chart repository"
	@echo "  helm-install    - Install the chart on K8s"
	@echo "  helm-upgrade    - Upgrade the release"
	@echo "  helm-uninstall  - Uninstall the release"
	@echo "  helm-status     - Show release status"
	@echo "  helm-template   - Render chart templates locally"
	@echo ""
	@echo "Override namespace: make helm-install NAMESPACE=my-ns"
