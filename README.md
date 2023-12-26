# ADDRESS SERVICE

## Microservice Description

The **Address Service** is a proof-of-concept microservice written in Go, utilizing the Gin framework and interacting with a PostgreSQL database. This microservice is designed to provide seamless CRUD (Create, Read, Update, Delete) operations for address entities. Whether you need to create new addresses, retrieve individual or multiple addresses, update existing records, or delete entries.

### Technologies Used:

- **Go Language:** The microservice is implemented in Go, leveraging its efficiency and performance benefits.

- **Gin Framework:** Built on top of Go, Gin provides a lightweight and fast HTTP web framework, making it ideal for building RESTful APIs.

- **PostgreSQL Database:** The microservice interacts with a PostgreSQL database to store and retrieve address data.

### Purpose:

This microservice serves as a proof of concept, demonstrating the capabilities of Go, Gin, and PostgreSQL in the context of a RESTful API for managing address data.

## Prerequisites

Before you start, ensure that you have [Golang](https://golang.org/dl/) installed on your machine. Follow the [official documentation](https://golang.org/doc/install) for installation instructions.

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/your-microservice.git
   cd your-microservice
   ```

2. Install dependencies:

```bash
go mod tidy
```

3. Start the project:

```bash
go run main.go
```

OR 

4. Build the project:

```bash
go build -o <any-name-you-like>

```

AND 

5.Execute the binary :

```bash
./the-name-you-picked
```

## Swagger Documentation

After starting the microservice locally, you can access the Swagger documentation at:

[Swagger Documentation](http://localhost:3001/swagger/index.html#/address/post_api_v1_addresses)
