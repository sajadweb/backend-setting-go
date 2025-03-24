## Go TCP microservice

This project is a **TCP-based microservice** built with **Golang**, designed to handle settings-related features such as category management.  
It exposes a custom TCP server that communicates with a NestJS client (or any other TCP client) over a predefined protocol.

## Features

- **TCP server implementation**
- **Pattern-based request handling (similar to routing)**
- **JSON message format with request/response IDs**
- **Category CRUD operations (Create & Find All)**
- **MongoDB integration for persistent storage**
- Simple and lightweight, ideal for microservice communication without HTTP overhead.

## Architecture Overview

```
NestJS Client ----> TCP Connection ----> Go TCP Server ----> MongoDB
```

- NestJS (or other clients) sends TCP requests with specific patterns like `create.category` or `find.all.categories`.
- Go TCP server listens, parses the request, and dispatches it to the appropriate handler (controller).
- Responses are sent back over the same TCP connection in JSON format.

---

## Project Structure

```
├── category
│   ├── controllers       // Handles incoming TCP requests for categories
│   ├── models            // Category data models (MongoDB schema)
│   ├── services          // Business logic and MongoDB operations
│   └── category.go       // Main entry to register category routes
├── config                // Environment loading (MongoDB URI, TCP server config)
├── microservices
│   └── tcp               // TCP server implementation and request handling
├── main.go               // Entry point of the application
└── README.md
```

---

## Request/Response Format

### Incoming TCP Request (JSON with header):

```
<length>#{
  "pattern": "create.category",
  "data": {
    "name": "Electronics",
    "icon": "📱"
  },
  "id": "123456"
}
```

### TCP Response:

```
<length>#{
  "isDisposed": true,
  "id": "123456",
  "response": {
    "data": { "_id": "abcdef123", "name": "Electronics", "icon": "📱" },
    "code": 200,
    "message": "Category created successfully",
    "error": false
  },
  "err": null
}
```

---

## Supported Patterns

| Pattern               | Description                     |
|----------------------|---------------------------------|
| `create.category`     | Creates a new category          |
| `find.all.categories` | Retrieves all categories        |

---

## Setup & Installation

1. **Clone the repository:**

```bash
git clone https://github.com/sajadweb/backend-setting-go.git
cd backend-setting-go
```

2. **Set environment variables:**

Create a `.env` file:

```bash
MONGODB_URI=mongodb://localhost:27017
MONGODB_DB=your_database
TCP_SERVER=0.0.0.0:9000
```

3. **Run the server:**

```bash
go mod tidy
go run main.go
```

---

## Connecting from NestJS

You can connect to this TCP server using NestJS's built-in microservices module:

```typescript
import { ClientProxyFactory, Transport } from '@nestjs/microservices';

const client = ClientProxyFactory.create({
  transport: Transport.TCP,
  options: { host: 'localhost', port: 9000 },
});
```

Then, you can send patterns like:

```typescript
client.send({ pattern: 'create.category', id: '123' }, { name: 'Electronics', icon: '📱' });
```

---

## Future Improvements

- Add more modules beyond category (e.g., user settings, preferences)
- Implement authentication for TCP connections
- Add unit tests
- Dockerize the service

---

## License

MIT License
