# Browsafe Backend

[![Go Report Card](https://goreportcard.com/badge/github.com/Sm69mu/Browsafe_backend)](https://goreportcard.com/report/github.com/Sm69mu/Browsafe_backend)
[![Go Version](https://img.shields.io/github/go-mod/go-version/Sm69mu/Browsafe_backend)](https://go.dev/)
[![License](https://img.shields.io/github/license/Sm69mu/Browsafe_backend)](LICENSE)

## Overview

Browsafe Backend is a Go-powered backend service that provides the server-side functionality for the Browsafe application. Built with modern Go practices and containerized with Docker for easy deployment.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [Docker Support](#docker-support)
- [API Documentation](#api-documentation)
- [Project Structure](#project-structure)
- [Contributing](#contributing)

## Features

- Written in Go (99.3% of codebase)
- Docker containerization support
- RESTful API architecture
- Scalable backend infrastructure

## Prerequisites

- Go 1.21 or higher
- Docker (for containerized deployment)
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Sm69mu/Browsafe_backend.git
cd Browsafe_backend
```

2. Install dependencies:
```bash
go mod download
```

## Running the Application

### Local Development

To run the application locally:

```bash
go run main.go
```

### Build

To build the application:

```bash
go build -o browsafe-backend
```

Then run the compiled binary:

```bash
./browsafe-backend
```

## Docker Support

The application includes Docker support (0.7% of codebase) for containerized deployment.

Build the Docker image:

```bash
docker build -t browsafe-backend .
```

Run the container:

```bash
docker run -p 8080:8080 browsafe-backend
```

## API Documentation

[API documentation to be added]

## Project Structure

```
browsafe-backend/
├── cmd/
│   └── main.go
├── internal/
│   ├── api/
│   ├── config/
│   ├── models/
│   └── services/
├── pkg/
├── Dockerfile
└── README.md
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the terms of the [LICENSE](LICENSE) file included in the repository.

## Contact

- GitHub: [@Sm69mu](https://github.com/Sm69mu)
