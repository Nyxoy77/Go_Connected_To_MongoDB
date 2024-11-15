# Go Connected to MongoDB

This project demonstrates how to connect a Go (Golang) backend to a MongoDB database. It provides a simple API for managing and interacting with MongoDB collections, focusing on basic CRUD operations. The project is intended for developers looking to understand how to set up and interact with MongoDB using Go.

## Table of Contents
- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

To get started, clone the repository and follow the instructions below to set up the necessary dependencies and configuration for MongoDB and Go.

## Prerequisites

- [Go](https://golang.org/dl/) installed (version 1.16+ recommended)
- MongoDB instance (either local or on the cloud)
- Git

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/Nyxoy77/Go_Connected_To_MongoDB.git
    cd Go_Connected_To_MongoDB
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Ensure MongoDB is running and accessible. This project requires a running MongoDB instance. You can install MongoDB locally or use a cloud service like MongoDB Atlas.

## Configuration

1. Create a `.env` file in the root directory and add your MongoDB URI and database information:

    ```plaintext
    MONGODB_URI="your_mongodb_uri"
    DATABASE_NAME="your_database_name"
    ```

2. Replace `your_mongodb_uri` and `your_database_name` with your MongoDB connection string and database name.

## Usage

1. Run the application:
    ```bash
    go run main.go
    ```

2. The server should now be running at `http://localhost:8080`.

## Endpoints

The following endpoints are available:

- **GET /api/data** - Fetches all records from the collection.
- **POST /api/data** - Adds a new record to the collection.
- **GET /api/data/{id}** - Fetches a single record by ID.
- **PUT /api/data/{id}** - Updates a record by ID.
- **DELETE /api/data/{id}** - Deletes a record by ID.

These endpoints are defined in the `main.go` file.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you have any improvements or fixes.

## License

This project is licensed under the MIT License.
