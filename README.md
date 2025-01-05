# Task Management System

This is a simple task management system built with Golang, featuring RESTful APIs, caching with Redis, and database persistence.

---

## Getting Started

### Prerequisites

Make sure you have the following installed:

- [Golang](https://golang.org/dl/) (version 1.23.1 or higher recommended)
- [Redis](https://redis.io/download)
- [PostgreSQL](https://www.postgresql.org/download) or your preferred database

---

### Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/aziemp66/dot-indonesia-technical-test
   cd dot-indonesia-technical-test
	 ```

2. **Create a .env File** 
  Use the provided example.env file as a template for creating your .env file. Replace the placeholder values with your desired configuration.

	```bash
	cp example.env .env
	```

3. **Edit the .env file**
	Edit the .env file value according to your needs

### Run the app

```bash
go run cmd/http/main.go
```
