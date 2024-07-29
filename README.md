# Go - HTTP Proxy Server

### This is HTTP proxy server for proxying requests to a third-party service that will accept requests from the client, send them to the specified services, and return responses to the client in JSON format. Also saving requests and responses locally. The server also provides a Swagger UI for API documentation.

## Getting Started

### Prerequisites:
    - Go 1.22 or later
    - Docker
### Installation:
1. Clone the repository:
   ```bash
   https://github.com/Aytya/proxy_server
   ```
2. Navigate into the project directory:
   ```bash
    cd proxy_server
   ```
3. Install dependencies:
   ```bash
    go get -u "github.com/swaggo/http-swagger"
    go get -u "github.com/go-chi/chi/v5/middleware"
    go get -u "github.com/go-chi/chi/v5"
   ```

##  Build and Run Locally:
 ### Build the application:
   ```bash
   make build
   ```
 ### Run the application:
   ```bash
   make run
   ```
 ### Generate Swagger Documentation:
   ```bash
   swag init -g cmd/main.go
   ```
## API Endpoints:
 ### Proxy Request:
 ```bash
    {
     "method": "GET",
     "url": "http://google.com",
     "headers": { "Authentication": "Basic bG9naW46cGFzc3dvcmQ=", "Authentications": "Basic bG9naW46cGFzc3dvcmQ=" }
    }
 ```
 ### Proxy Response:
 ```bash
   {
    "id": "20240729133515",
    "status": 301,
    "headers": {
        "Cache-Control": [
            "public, max-age=2592000"
        ],
        "Content-Length": [
            "219"
        ],
        "Content-Security-Policy-Report-Only": [
            "object-src 'none';base-uri 'self';script-src 'nonce-VtQPKVi6nhCLiMooge89Cw' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp"
        ],
        "Content-Type": [
            "text/html; charset=UTF-8"
        ],
        "Date": [
            "Mon, 29 Jul 2024 13:35:15 GMT"
        ],
        "Expires": [
            "Wed, 28 Aug 2024 13:35:15 GMT"
        ],
        "Location": [
            "http://www.google.com/"
        ],
        "Server": [
            "gws"
        ],
        "X-Frame-Options": [
            "SAMEORIGIN"
        ],
        "X-Xss-Protection": [
            "0"
        ]
    },
    "length": 219
  }
 ```
 ### Swagger Documentation
   - URL: /swagger/
