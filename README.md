## README.md (English)

# Go Gin Slog Cleanenv Template

This is a starter template for building Go web applications using:

* **Gin:** A fast and lightweight web framework for Go.
* **Slog:** Go's structured logging package.
* **Cleanenv:** A library for reading configuration from environment variables, files, and command-line flags.

## Features:

* Pre-configured logging with slog and a custom handler for pretty-printed logs.
* Environment-based configuration loading using cleanenv.
* Basic project structure for organizing code (handlers, services, routes).
* Dockerfile for containerization.
* Docker Compose for running the application with dependencies.

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/go-gin-slog-cleanenv-template.git
   ```
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Build and run:
   ```bash
   go run cmd/server/main.go
   ```
4. Access the application at `http://localhost:9090`.

## Configuration

The application configuration is read from the following sources in order of precedence:

1. Command-line flag `-config`: Specifies the path to the configuration file.
2. Environment variable `CONFIG_PATH`: Specifies the path to the configuration file.
3. Default configuration file `config/server/config.yaml`.

## Docker

You can build and run the application using Docker:

```bash
docker-compose up --build
```

This will build the Docker image and start the application in a container. The application will be accessible at `http://localhost:9090`.



