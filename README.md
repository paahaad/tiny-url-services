# URL Shortener Service

## Features

- **URL Shortening:** Convert long URLs into short, easy-to-share links.
- **Custom Aliases:** Option to create custom short URLs.
- **Redirects:** Automatically redirect to the original URL when the short link is accessed.
- **Analytics:** Track the number of clicks on each shortened URL (optional).
- **Expiration:** Set an expiration date for short URLs (optional).

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16+)
- A PostgreSQL database (or another supported database)
- [Docker](https://docs.docker.com/get-docker/) (optional, for containerization)

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/url-shortener-go.git
   cd url-shortener-go
   ```

2. **Set up environment variables:**

   Create a `.env` file in the root directory:

   ```bash
   touch .env
   ```

   Add the following variables:

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_username
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   SERVER_PORT=8080
   ```

3. **Install dependencies:**

   ```bash
   go mod tidy
   ```

4. **Run database migrations:**

   You can use [golang-migrate](https://github.com/golang-migrate/migrate) or any other migration tool:

   ```bash
   migrate -database postgres://$DB_USER:$DB_PASSWORD@localhost:$DB_PORT/$DB_NAME?sslmode=disable -path migrations up
   ```

5. **Run the application:**

   ```bash
   go run main.go
   ```

   The service will be available at `http://localhost:8080`.

## Usage

1. **Shorten a URL:**

   Send a `POST` request to `/api/v1/shorten` with the following JSON payload:

   ```json
   {
     "url": "https://www.example.com",
     "customAlias": "example" // Optional
   }
   ```

   Example using `curl`:

   ```bash
   curl -X POST http://localhost:8080/api/v1/shorten -H "Content-Type: application/json" -d '{"url":"https://www.example.com","customAlias":"example"}'
   ```

2. **Access a shortened URL:**

   Visit `http://localhost:8080/{shortCode}` in your browser.

3. **Analytics (optional):**

   If enabled, access analytics at `/api/v1/analytics/{shortCode}`.

## Running with Docker

1. **Build the Docker image:**

   ```bash
   docker build -t url-shortener-go .
   ```

2. **Run the container:**

   ```bash
   docker run --env-file .env -p 8080:8080 url-shortener-go
   ```

## Testing

Run unit tests with the following command:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes.