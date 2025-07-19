# Go API - Product Management

A REST API built with Go (Golang) for product management, featuring a clean architecture with PostgreSQL database integration.

## 🏗️ Architecture

This project follows clean architecture principles with the following layers:

- **Controller**: HTTP request handlers
- **UseCase**: Business logic layer
- **Repository**: Data access layer
- **Model**: Data structures and entities
- **Database**: PostgreSQL connection management

## 📁 Project Structure

```
├── cmd/
│   └── main.go              # Application entry point
├── controller/
│   └── product_controller.go # HTTP handlers
├── db/
│   └── conn.go              # Database connection
├── model/
│   ├── product.go           # Product entity
│   └── response.go          # Response models
├── repository/
│   └── product_repository.go # Data access layer
├── usecase/
│   └── product_usecase.go   # Business logic
├── docker-compose.yml       # Docker services configuration
├── dockerfile              # Docker image configuration
├── go.mod                  # Go module dependencies
└── go.sum                  # Dependency checksums
```

## 🚀 Features

- **Product Management**: Create, read, and list products
- **RESTful API**: Clean REST endpoints
- **PostgreSQL Integration**: Persistent data storage
- **Docker Support**: Containerized application and database
- **Health Check**: API health monitoring endpoint
- **Clean Architecture**: Separation of concerns with dependency injection

## 📋 Prerequisites

- Go 1.24.5 or higher
- Docker and Docker Compose
- PostgreSQL (if running without Docker)

## 🛠️ Installation

### Using Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd Golang
   ```

2. **Start the application with Docker Compose**
   ```bash
   docker-compose up --build
   ```

   This will start:
   - Go API on port `8000`
   - PostgreSQL database on port `5432`

### Manual Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd Golang
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up PostgreSQL database**
   - Create a database named `go_products`
   - Update database connection details in `db/conn.go` if needed

4. **Run the application**
   ```bash
   go run cmd/main.go
   ```

## 🔗 API Endpoints

### Health Check
- **GET** `/health` - Check API health status

### Products
- **GET** `/products` - List all products
- **POST** `/products` - Create a new product
- **GET** `/products/:id` - Get a product by ID

## 📝 API Usage Examples

### Health Check
```bash
curl http://localhost:8000/health
```

### Create a Product
```bash
curl -X POST http://localhost:8000/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop",
    "price": 999.99
  }'
```

### Get All Products
```bash
curl http://localhost:8000/products
```

### Get Product by ID
```bash
curl http://localhost:8000/products/1
```

## 🗃️ Database Configuration

### Docker Environment
- **Host**: `go_db` (container name)
- **Port**: `5432`
- **Database**: `go_products`
- **Username**: `postgres`
- **Password**: `1234`

### Local Environment
Update the database connection in `db/conn.go` according to your local PostgreSQL setup.

## 🔧 Technologies Used

- **[Go](https://golang.org/)** - Programming language
- **[Gin](https://gin-gonic.com/)** - HTTP web framework
- **[PostgreSQL](https://www.postgresql.org/)** - Database
- **[lib/pq](https://github.com/lib/pq)** - PostgreSQL driver
- **[Docker](https://www.docker.com/)** - Containerization

## 📦 Dependencies

Key dependencies include:
- `github.com/gin-gonic/gin` - Web framework
- `github.com/lib/pq` - PostgreSQL driver
- `github.com/go-playground/validator/v10` - Input validation

## 🐳 Docker Services

The `docker-compose.yml` defines two services:

1. **go-app**: The API application
   - Built from local Dockerfile
   - Exposes port 8000
   - Depends on the database service

2. **go_db**: PostgreSQL database
   - Uses PostgreSQL 15 image
   - Persistent data with named volume
   - Pre-configured with database credentials

## 🚦 Getting Started

1. Ensure Docker and Docker Compose are installed
2. Run `docker-compose up --build`
3. The API will be available at `http://localhost:8000`
4. Test the health endpoint: `curl http://localhost:8000/health`

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 👨‍💻 Author

Antonio Carlos Santos Junior

---

**Note**: This project was created as part of a YouTube tutorial series for learning Go API development.
