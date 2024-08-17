# URL Shortener

A simple URL shortener application built with Golang, Postgres, and Redis. This app takes a long URL provided by a user and returns a shortened version. The shortened URLs can be stored and retrieved efficiently using Redis for quick lookups and Postgres for persistence.

## Architecture diagram

![arch](/docs/image.png)

## Features- Shorten long URLs

- Store and retrieve URLs using Postgres and Redis
- Redirect users to the original URL when the shortened link is accessed

## Tech Stack-**Backend**: Golang

-**Database**: Postgres (for persistence) -**Cache**: Redis (for fast lookups)

## Setup

1. Install dependencies:

```bash
go mod tidy
```

2. Set up environment variables:
   Create a `.env` file from `.env.exampl` and replace the values with your secrets

```bash
cp .env.example .env
```

3. Set up docker container

```bash
docker-compose up -d
```

4. Run the app

```bash
go run cmd/main.go
```

5. Open browser at `http://localhost:8080`

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
