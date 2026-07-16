<h1 align="center">Golang API CRUD</h1>
<p align="center">
<i>A simple REST API built with Go, featuring a complete CI/CD pipeline and automated Kubernetes deployment.</i>
</p>

---

## Table of Contents

* [Overview](#overview)
* [Solution Architecture](#solution-architecture)
* [Code Architecture](#code-architecture)
* [Technology Stack](#technology-stack)
* [Infrastructure](#infrastructure)
* [Local Setup](#local-setup)
* [Docker](#docker)
* [Endpoints](#endpoints)
* [CI/CD Pipeline](#cicd-pipeline)
* [Security (DevSecOps)](#security-devsecops)
* [Workflow Diagram](#workflow-diagram)
* [Next Repository](#next-repository)
* [Contact](#contact)

---

## Overview

This project is a REST API developed in Go that performs CRUD operations while persisting data in PostgreSQL. The application follows a layered architecture (inspired by Clean Architecture) and includes a complete CI/CD pipeline covering automated testing, code quality analysis, security scanning, and automated deployments to Kubernetes environments (Development and Production).

The primary goal of this project is to demonstrate end-to-end software development and delivery best practices using a DevSecOps approach from writing the code to deploying it into production, including automated quality and security validation throughout the pipeline.

---

## Solution Architecture

```text
Developer
     │
     ▼
GitHub
     │
     ▼
GitHub Actions
     │
     ├── Tests
     ├── SonarCloud
     ├── Docker Build
     ├── Trivy Scan
     ├── Push to GHCR
     └── Kubernetes Deployment
                     │
                     ▼
            Kubernetes Cluster
                     │
                     ▼
              PostgreSQL
```

---

## Code Architecture

The project is organized into layers with well-defined responsibilities.

```text
├── cmd/          # Application entry point (main.go)
├── controller/   # HTTP layer - handles requests and responses
├── usecase/      # Business logic
├── repository/   # Data access layer / database abstraction
├── model/        # Entities and data structures
├── db/           # Database connection and configuration
├── migrations/   # PostgreSQL schema migrations
```

---

## Technology Stack

### Application Stack

* Go (Golang)
* Gin
* PostgreSQL

### DevOps Stack

* Docker
* Kubernetes
* GitHub Actions
* GitHub Container Registry (GHCR)
* Kustomize
* SonarCloud
* Trivy

---

## Infrastructure

The project uses a containerized architecture orchestrated by Kubernetes.

* Docker for application packaging
* PostgreSQL as the relational database
* Kubernetes for container orchestration
* Kustomize for environment management
* GitHub Container Registry (GHCR) for image storage
* GitHub Actions for CI/CD automation

---

## Local Setup

```bash
# Clone the repository
git clone https://github.com/alexsandroocanha/Golang-Api-Crud.git
cd Golang-Api-Crud

# Configure environment variables
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=golang_crud

# Run the application
go run cmd/main.go
```

The application runs on port **8000** by default.

---

## Docker

```bash
docker build -t golang-api-crud .

docker run \
-p 8000:8000 \
--env-file .env \
golang-api-crud
```

---

## Endpoints

| Method | Endpoint              | Description              |
| ------ | --------------------- | ------------------------ |
| GET    | `/ping`               | Health Check             |
| GET    | `/products`           | List all products        |
| GET    | `/product/:productId` | Retrieve a product by ID |
| POST   | `/product`            | Create a new product     |

---

<h1 align="center">CI/CD Pipeline</h1>
<p align="center">
<i>Build Once, Promote Many</i>
</p>

This project follows the **Build Once, Promote Many** deployment strategy.

The Docker image is built only once after a merge into the `dev` branch.

Once all automated tests and security validations have passed, the same image is promoted to production by applying semantic version tags, ensuring consistency across every environment.

Kubernetes manifests are maintained in a separate repository using **Kustomize** with dedicated overlays for both **Development** and **Production** environments.

### Workflows

| Workflow                 | Purpose                                              |
| ------------------------ | ---------------------------------------------------- |
| Development Pull Request | Tests + Code Coverage + SonarCloud                   |
| Development Deployment   | Build + Push + Trivy Scan + Development Deployment   |
| Production Release       | Image Promotion + Versioning + Production Deployment |

---

## Security (DevSecOps)

The CI/CD pipeline includes multiple automated security and quality controls.

* Automated tests
* Code coverage
* SonarCloud analysis
* Trivy image scanning
* GitHub Container Registry
* GitHub Actions pinned by SHA (Supply Chain Security)
* Deployment gates
* Build Once, Promote Many strategy

---

## Workflow Diagram

```text
Pull Request
 │
 ▼
Tests
 │
 ▼
SonarCloud
 │
 ▼
Docker Build
 │
 ▼
Push to GHCR
 │
 ▼
Trivy Scan
 │
 ▼
Development Deployment
 │
 ▼
Merge to Main
 │
 ▼
Image Promotion
 │
 ▼
Production Deployment
```

---

## Next Repository

This repository contains the application source code.

The Kubernetes infrastructure is maintained in a separate manifests repository.

[![GitHub Pages](https://img.shields.io/badge/Golang%20API%20Crud%20Manifests-121013?style=for-the-badge\&logo=github\&logoColor=white)](https://github.com/alexsandroocanha/Golang-Api-Crud-Manifests)

---

## Contact

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge\&logo=linkedin\&logoColor=white)](https://www.linkedin.com/in/alexsandro-ocanha-rodrigues-77149a35b/)

[![Instagram](https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge\&logo=instagram\&logoColor=white)](https://www.instagram.com/alexsandro.pcap/)

[![Gmail](https://img.shields.io/badge/Gmail-D14836?style=for-the-badge\&logo=gmail\&logoColor=white)](mailto:alexsandroocanha@gmail.com)
