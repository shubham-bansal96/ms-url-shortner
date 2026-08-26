# ms-url-shortner — URL Shortener Microservice

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat&logo=go)
![Gin](https://img.shields.io/badge/Gin-v1.12-blue?style=flat)
![Prometheus](https://img.shields.io/badge/Prometheus-Metrics-E6522C?style=flat&logo=prometheus)
![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?style=flat&logo=docker)
![Helm](https://img.shields.io/badge/Helm-K8s_Deploy-0F1689?style=flat&logo=helm)

A lightweight **URL Shortener** microservice built in Go using the [Gin](https://github.com/gin-gonic/gin) framework. Accepts long URLs via a REST API and returns shortened URLs with an 8-character unique identifier. Includes Prometheus metrics, rate limiting, structured logging, and is deployable to Kubernetes via Helm.

---

## Features

- **URL Shortening** — generates short URLs with 8-character unique IDs
- **In-Memory Cache** — previously shortened URLs are stored in memory for instant retrieval on repeat requests
- **Input Validation** — rejects empty URLs and those missing `http://` or `https://` prefix
- **Smart Length Check** — URLs of 20 characters or fewer are returned as-is (already short enough)
- **Prometheus Metrics** — built-in `/metrics` endpoint for monitoring request counts, latencies, and more
- **Rate Limiting** — configurable rate limiter middleware to protect against abuse
- **Structured Logging** — Logrus-based logging with configurable log levels
- **Profiling** — pprof endpoints enabled for runtime performance analysis
- **Kubernetes Ready** — Helm chart for easy deployment and scaling on K8s
- **Swagger Documentation** — auto-generated API spec from source code annotations

---

## Architecture

```
┌─────────────────────────────────────────────────────────────────────┐
│                         REQUEST FLOW                                  │
├─────────────────────────────────────────────────────────────────────┤
│                                                                      │
│  Client ──→ Gin Router ──→ Rate Limiter ──→ Metrics Middleware       │
│                                                      │               │
│                                                      ▼               │
│                                              Base Controller         │
│                                                      │               │
│                                                      ▼               │
│                                           URL Shortener Service      │
│                                              │            │          │
│                                              ▼            ▼          │
│                                        Validate URL   Check Cache    │
│                                              │            │          │
│                                              ▼            ▼          │
│                                        Generate UID   Return Cached  │
│                                              │                       │
│                                              ▼                       │
│                                        Store in Map                  │
│                                              │                       │
│                                              ▼                       │
│                                        JSON Response                 │
│                                                                      │
└─────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────┐
│                       OBSERVABILITY                                   │
├─────────────────────────────────────────────────────────────────────┤
│                                                                      │
│  /metrics ──→ Prometheus ──→ Grafana Dashboard                       │
│  /debug/pprof ──→ CPU / Memory / Goroutine Profiling                 │
│                                                                      │
└─────────────────────────────────────────────────────────────────────┘
```

---

## Tech Stack

| Component | Technology | Purpose |
|-----------|-----------|---------|
| Language | Go 1.26 | Core application |
| Framework | Gin v1.12 | HTTP router and middleware |
| Metrics | Prometheus client_golang | Request metrics and monitoring |
| Logging | Logrus | Structured logging with levels |
| Profiling | pprof | Runtime performance analysis |
| Rate Limiting | golang.org/x/time | Token bucket rate limiter |
| Containerization | Docker (multi-stage) | Lightweight Alpine-based image |
| Orchestration | Helm + Kubernetes | Deployment, scaling, and management |
| API Docs | go-swagger | Auto-generated Swagger spec |

---

## Project Structure

```
ms-url-shortner/
├── main.go                          # Application entry point
├── config.yml                       # Application configuration
├── Dockerfile                       # Multi-stage Docker build
├── Makefile                         # Build, test, and deploy shortcuts
├── swagger.yaml                     # Generated API specification
├── app/
│   ├── config/
│   │   ├── config.go                # Configuration loader (YAML)
│   │   └── constant.go              # Config constants
│   ├── controller/
│   │   ├── base_controller.go       # HTTP handlers (ping, shorten)
│   │   └── base_controller_test.go  # Controller unit tests
│   ├── model/
│   │   ├── url_shortner.go          # Request/domain models
│   │   ├── url_shortner_test.go     # Model unit tests
│   │   └── responseDTO.go           # Response DTOs
│   ├── services/
│   │   ├── url_shortner_service.go      # URL shortening business logic
│   │   ├── url_shortner_service_test.go # Service unit tests
│   │   └── metrics.go                   # Prometheus metric definitions
│   ├── middleware/
│   │   ├── metrics.go               # Prometheus metrics middleware
│   │   └── rate_limiter.go          # Rate limiting middleware
│   ├── route/
│   │   ├── route.go                 # Route registration
│   │   └── constant.go              # Route path constants
│   ├── logging/
│   │   └── logger.go                # Logrus logger setup
│   ├── utils/
│   │   └── json-response.go         # JSON response helpers
│   ├── docs/
│   │   └── docs.go                  # Swagger annotations
│   └── test-helper/
│       ├── config_mock.go           # Config mocks for tests
│       └── service_mock.go          # Service mocks for tests
└── url-shortner-charts/             # Helm chart for K8s deployment
    ├── Chart.yaml
    ├── values.yaml
    └── templates/
```

---

## Prerequisites

### Go 1.26+

```bash
# macOS
brew install go

# Verify
go version
```

### Docker (for containerized runs)

```bash
# macOS
brew install --cask docker

# Verify
docker --version
```

### Helm & Kubernetes (for K8s deployment)

```bash
# macOS
brew install helm kubectl

# Verify
helm version
kubectl version --client
```

---

## Quick Start

```bash
# 1. Clone the repository
git clone https://github.com/shubham-bansal96/ms-url-shortner.git
cd ms-url-shortner

# 2. Run locally
make run

# 3. Test the health endpoint
curl http://localhost:4242/ms-url-shortner/ping

# 4. Shorten a URL
curl -X POST http://localhost:4242/ms-url-shortner/getshorturl \
  -H "Content-Type: application/json" \
  -d '{"url": "https://www.example.com/very/long/path/to/resource"}'
```

Or using Docker:

```bash
make docker-run
```

Or deploy to Kubernetes:

```bash
make helm-install
```

---

## API Reference

### Base URL

```
http://localhost:4242/ms-url-shortner
```

### Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/ping` | GET | Health check — verify the service is running |
| `/getshorturl` | POST | Shorten a URL |
| `/metrics` | GET | Prometheus metrics |
| `/debug/pprof` | GET | Profiling endpoints |

### POST `/getshorturl`

**Request:**

```json
{
  "url": "https://www.example.com/very/long/path/to/resource"
}
```

**Success Response (200):**

```json
{
  "data": {
    "url": "https://shorturl.com/ecfd35c4"
  },
  "error": null
}
```

**Error Responses:**

| Code | Condition | Example |
|------|-----------|---------|
| 400 | Malformed JSON body | Request body cannot be parsed |
| 422 | Invalid or empty URL | Missing `http://` or `https://` prefix |

```json
{
  "data": null,
  "error": {
    "code": 422,
    "message": "invalid url"
  }
}
```

---

## URL Shortening Rules

1. **Already short** — if the URL is 20 characters or fewer, it is returned as-is
2. **Invalid URL** — empty strings or URLs without `http://` / `https://` prefix are rejected with a 422 error
3. **Shortened URL** — valid long URLs get an 8-character unique ID: `https://shorturl.com/<8-char-uid>`
4. **Cached** — once shortened, the mapping is stored in memory; repeat requests return the same short URL instantly

---

## Configuration

The application is configured via `config.yml`:

| Key | Default | Description |
|-----|---------|-------------|
| `Logging.LogLevel` | `debug` | Log level (debug, info, warn, error, fatal) |
| `MSName` | `ms-url-shortner` | Microservice name (used as route prefix) |
| `Environment` | `dev` | Environment identifier |

---

## Running Tests

```bash
# Run all tests
make test

# Run with verbose output
make test-verbose

# Run with coverage report
make test-coverage
```

---

## Docker

The image uses a multi-stage build (builder on `golang:alpine`, runtime on `alpine:latest`) for a minimal footprint.

```bash
# Build the image
make docker-build

# Build and run (exposes port 4000)
make docker-run

# Or pull from Docker Hub directly
docker run -td -p 4242:4000 shubhambansal96/msurlshortner
```

**Docker Hub:** [shubhambansal96/msurlshortner](https://hub.docker.com/r/shubhambansal96/msurlshortner)

---

## Kubernetes Deployment (Helm)

A Helm chart is available for deploying to any Kubernetes cluster.

**Chart Repository:** https://shubham-bansal96.github.io/ms-url-shortner/

```bash
# Add the Helm repo
make helm-repo-add

# Install the chart
make helm-install

# Check status
make helm-status

# Upgrade after chart changes
make helm-upgrade

# Uninstall
make helm-uninstall

# Override namespace
make helm-install NAMESPACE=production
```

### Helm Values (defaults)

| Value | Default | Description |
|-------|---------|-------------|
| `replicaCount` | `2` | Number of pod replicas |
| `image.repository` | `shubhambansal96/msurlshortner` | Docker image |
| `image.tag` | `latest` | Image tag |
| `service.type` | `NodePort` | Kubernetes service type |
| `service.port` | `4242` | Service port |
| `prometheus.enabled` | `false` | Enable ServiceMonitor for Prometheus |
| `autoscaling.enabled` | `false` | Enable HPA |

---

## Swagger / API Docs

```bash
# Generate the spec from source code annotations
make swagger

# Serve the Swagger UI in browser
make serve-swagger
```

Alternatively, paste `swagger.yaml` into [Swagger Editor](https://editor.swagger.io/).

> **Install go-swagger:** https://goswagger.io/install.html

---

## Makefile Commands

```bash
make help
```

| Target | Description |
|--------|-------------|
| `build` | Build the binary |
| `run` | Run the application |
| `test` | Run tests |
| `test-verbose` | Run tests with verbose output |
| `test-coverage` | Run tests with coverage report |
| `fmt` | Format code |
| `vet` | Run go vet |
| `lint` | Run golangci-lint |
| `vendor` | Update vendor directory |
| `tidy` | Tidy go modules |
| `clean` | Remove build artifacts |
| `docker-build` | Build Docker image |
| `docker-run` | Build and run Docker container |
| `swagger` | Generate swagger spec |
| `serve-swagger` | Serve swagger UI |
| `helm-repo-add` | Add Helm chart repository |
| `helm-install` | Install the chart on K8s |
| `helm-upgrade` | Upgrade the release |
| `helm-uninstall` | Uninstall the release |
| `helm-status` | Show release status |
| `helm-template` | Render chart templates locally |

---

## Acknowledgments

- [Gin](https://github.com/gin-gonic/gin) — High-performance HTTP framework for Go
- [Prometheus](https://prometheus.io/) — Monitoring and alerting toolkit
- [Logrus](https://github.com/sirupsen/logrus) — Structured logger for Go
- [go-swagger](https://goswagger.io/) — Swagger 2.0 implementation for Go
- [Helm](https://helm.sh/) — Kubernetes package manager
