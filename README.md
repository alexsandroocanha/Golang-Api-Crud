<h1 align="center">Golang Api Crud</h1>
<p align="center">
<i>API REST simples desenvolvida em Go com pipeline CI/CD completa e deploy automatizado em Kubernetes.</i>
</p>

---

# Sumário

- [Visão Geral](#visão-geral)
- [Arquitetura da Solução](#arquitetura-da-solução)
- [Arquitetura do Código](#arquitetura-do-código)
- [Stack Tecnológica](#stack-tecnológica)
- [Infraestrutura](#infraestrutura)
- [Execução Local](#execução-local)
- [Docker](#docker)
- [Endpoints](#endpoints)
- [Pipeline CI/CD](#pipeline-cicd)
- [Segurança (DevSecOps)](#segurança-devsecops)
- [Diagrama do Fluxo](#diagrama-do-fluxo)
- [Contato](#contato)

---

# Visão Geral

API REST desenvolvida em Go, com operações de CRUD, persistindo dados em PostgreSQL. O projeto segue uma arquitetura em camadas (inspirada em Clean Architecture) e conta com uma pipeline de CI/CD completa, cobrindo testes, qualidade de código, segurança e deploy automatizado em Kubernetes (Dev e Prod).

O objetivo deste projeto é colocar em prática, de ponta a ponta, boas práticas de desenvolvimento e entrega de software utilizando práticas de DevSecOps. Desde a escrita do código até o deploy em produção, passando por validações automatizadas de qualidade e segurança.

---

# Arquitetura da Solução

```
Developer
     │
     ▼
GitHub
     │
     ▼
GitHub Actions
     │
     ├── Testes
     ├── SonarCloud
     ├── Build Docker
     ├── Trivy Scan
     ├── Push GHCR
     └── Deploy Kubernetes
                     │
                     ▼
            Kubernetes Cluster
                     │
                     ▼
              PostgreSQL
```

---

# Arquitetura do Código

O projeto é organizado em camadas com responsabilidades bem definidas:

```text
├── cmd/          # Ponto de entrada da aplicação (main.go)
├── controller/   # Camada HTTP - recebe requisições e trata respostas
├── usecase/      # Regras de negócio da aplicação
├── repository/   # Acesso a dados / abstração do banco
├── model/        # Entidades e estruturas de dados
├── db/           # Conexão e configuração do banco de dados
├── migrations/   # Versionamento do schema do PostgreSQL
```

---

# Stack Tecnológica

## Application Stack

- Go (Golang)
- Gin
- PostgreSQL

## DevOps Stack

- Docker
- Kubernetes
- GitHub Actions
- GitHub Container Registry (GHCR)
- Kustomize
- SonarCloud
- Trivy

---

# Infraestrutura

O projeto utiliza uma arquitetura baseada em containers e Kubernetes.

- Docker para empacotamento da aplicação
- PostgreSQL como banco de dados
- Kubernetes para orquestração
- Kustomize para gerenciamento dos ambientes
- GitHub Container Registry para armazenamento das imagens
- GitHub Actions para CI/CD

---

# Execução Local

```bash
# clonar o repositório
git clone https://github.com/alexsandroocanha/Golang-Api-Crud.git
cd Golang-Api-Crud

# configurar variáveis de ambiente
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=golang_crud

# rodar a aplicação
go run cmd/main.go
```

A aplicação sobe por padrão na porta **8000**.

---

# Docker

```bash
docker build -t golang-api-crud .

docker run \
-p 8000:8000 \
--env-file .env \
golang-api-crud
```

---

# Endpoints

| Método | Endpoint | Descrição |
|---------|----------|-----------|
| GET | `/ping` | Health Check |
| GET | `/products` | Lista todos os produtos |
| GET | `/product/:productId` | Busca produto por ID |
| POST | `/product` | Cria um novo produto |

---

<h1 align="center">Pipeline CI/CD</h1>
<p align="center">
<i>Build Once, Promote Many</i>
</p>

A estratégia utilizada neste projeto é **Build Once, Promote Many**.

A imagem Docker é construída apenas uma vez quando ocorre merge na branch `dev`.

Após todos os testes e validações de segurança, a mesma imagem é promovida para produção utilizando apenas novas tags semânticas, eliminando diferenças entre os ambientes.

Os manifests Kubernetes ficam em um repositório separado utilizando **Kustomize** com overlays para **Dev** e **Prod**.

## Workflows

| Workflow | Objetivo |
|----------|----------|
| Development Pull Request | Testes + Cobertura + SonarCloud |
| Development Deploy | Build + Push + Trivy + Deploy Dev |
| Production Release | Promote Image + Versionamento + Deploy Prod |

---

# Segurança (DevSecOps)

Durante a pipeline são aplicadas diversas validações automáticas.

- Testes automatizados
- Cobertura de código
- SonarCloud
- Trivy Image Scan
- GitHub Container Registry
- Actions fixadas por SHA (Supply Chain Security)
- Gates de Deploy
- Build Once Promote Many

---

# Diagrama do Fluxo

```text
PR
 │
 ▼
Testes
 │
 ▼
SonarCloud
 │
 ▼
Build Docker
 │
 ▼
Push GHCR
 │
 ▼
Trivy Scan
 │
 ▼
Deploy Dev
 │
 ▼
Merge Main
 │
 ▼
Promote Image
 │
 ▼
Deploy Produção
```
---

## Próximo Repositório

Este repositório contém a aplicação.

A infraestrutura Kubernetes encontra-se no repositório de manifests.

[![GitHub Pages](https://img.shields.io/badge/Golang%20API%20Crud%20Manifests-121013?style=for-the-badge&logo=github&logoColor=white)](https://github.com/alexsandroocanha/Golang-Api-Crud-Manifests)

---

# Contato

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/alexsandro-ocanha-rodrigues-77149a35b/)

[![Instagram](https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge&logo=instagram&logoColor=white)](https://www.instagram.com/alexsandro.pcap/)

[![Gmail](https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:alexsandroocanha@gmail.com)

