# Golang Starter Kit - Restful API with Gin and Gorm
> A starter kit for building RESTful APIs in Golang using Gin and Gorm.
This repository provides a basic structure and setup for creating RESTful APIs in Go, utilizing the Gin web framework and Gorm ORM for database interactions.
- **Gin**: A web framework written in Go that is known for its speed and small memory footprint.
- **Gorm**: An ORM library for Go that provides a simple and powerful way to interact with databases.

## Features
- RESTful API structure
- Middleware support
- Database migrations
- Error handling
- Logging
- Environment configuration
- Dependency injection

## Getting Started
### Prerequisites
- Go 1.18 or later
- PostgreSQL or any other supported database
- Docker (optional, for running the database in a container)

### Installation
1. Clone the repository:
   ```bash
   git clone repository_url
   cd repository_name
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Set up your environment variables. You can create a `.env` file in the root directory with the following content:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   PORT=8080
   ```
4. Run the application:
   ```bash
   go run main.go
   ```

### Running with Docker
If you prefer to run the application with Docker, you can use the provided `docker-compose.yml` file. Make sure you have Docker and Docker Compose installed.
1. Build and run the containers:
   ```bash
   docker-compose up --build
   ```
2. Access the application at `http://localhost:8080`.

### Make a Tests with Gomock
To ensure the application works as expected, you can run the provided tests. The tests are located in the `tests` directory and cover various aspects of the API functionality.
1. Make mocks for the tests e.g:
   ```bash
   mockgen -source=internal/repository/auth.repository.go -destination=mocks/services/auth.repository.mock.go -package=mocks
   ```
2. Then make tests file for example:
   ```go
   package tests

   import (
       "testing"
       "github.com/stretchr/testify/assert"
       "github.com/yourusername/yourproject/internal/repository/mocks"
       "github.com/golang/mock/gomock"
   )

   func TestExample(t *testing.T) {
       ctrl := gomock.NewController(t)
       defer ctrl.Finish()

       mockRepo := mocks.NewMockAuthRepository(ctrl)
       // Set expectations and call methods on mockRepo
       assert.NotNil(t, mockRepo)
   }
   ```

### Running Tests
To run the tests, you can use the following command:
```bash
go test ./...
```

### Contributing
Contributions are welcome! If you have suggestions for improvements or new features, feel free to open an issue or submit a pull request.

### License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.