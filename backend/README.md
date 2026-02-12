# Backend - ShuttleSpark Lab

This is the backend service for ShuttleSpark Lab, built using the [Huma](https://huma.rocks) framework with the [Chi](https://go-chi.io) router.

## Getting Started

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Run the server:
   ```bash
   go run main.go
   ```

3. Visit the API documentation:
   Open [http://localhost:8888/docs](http://localhost:8888/docs) in your browser.

## API Endpoints

- `GET /greeting/{name}`: Get a greeting message.
