# URL Shortener API

This is a simple URL shortener API built using Go. It allows users to shorten URLs and redirect to the original URLs using the shortened version.

## Features

- Shorten a given URL.
- Redirect to the original URL using the shortened URL.
- Simple in-memory storage for URLs.

## Endpoints

### 1. Shorten URL

- **Endpoint**: `/getUrl`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "url": "https://example.com"
  }
  ```
- **Response**:
  ```json
  {
    "shortUrl": "abc12345"
  }
  ```

### 2. Redirect to Original URL

- **Endpoint**: `/redirect/{shortUrl}`
- **Method**: `GET`
- **Description**: Redirects to the original URL associated with the given `shortUrl`.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/url-shortener.git
   ```
2. Navigate to the project directory:
   ```bash
   cd url-shortener
   ```

### Running the Server

1. Run the server:
   ```bash
   go run main.go
   ```
2. The server will start on port 3000.

### Usage

- Use a tool like `curl` or Postman to interact with the API.
- Example using `curl` to shorten a URL:
  ```bash
  curl -X POST -H "Content-Type: application/json" -d '{"url": "https://example.com"}' http://localhost:3000/getUrl
  ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by various URL shortener services.
