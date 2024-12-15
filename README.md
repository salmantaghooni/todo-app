# Todo App

A Todo application built with Go, Gin, GORM, AWS S3, SQS, Docker, and Docker Compose. It allows users to upload files and create Todo items with optional file attachments.

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)
- [Go](https://golang.org/dl/) (version 1.23 or newer)

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/todo-app.git
cd todo-app
```

### 2. Create a `.env` File

Create a `.env` file in the root directory with the following content:

```env
SERVER_PORT=8080

DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=todo_db
DB_SSLMODE=disable

S3_ENDPOINT=http://localstack:4566
S3_REGION=us-east-1
S3_ACCESS_KEY_ID=test
S3_SECRET_ACCESS_KEY=test
S3_BUCKET=uploads

SQS_ENDPOINT=http://localstack:4566
SQS_REGION=us-east-1
SQS_QUEUE_URL=http://localstack:4566/000000000000/todo-queue
```

### 3. Start the Application

Run the application using Make:

```bash
make run
```

This command will:

- Start PostgreSQL and LocalStack services using Docker Compose.
- Build and run the Go application.

### 4. Apply Database Migrations

The application uses GORM's `AutoMigrate` to apply migrations automatically on startup. If you prefer using raw SQL migrations, you can integrate a migration tool like `golang-migrate`.

### 5. Accessing the API

The API will be available at `http://localhost:8080`.

## API Endpoints

### 1. Upload File

- **Endpoint:** `/upload`
- **Method:** `POST`
- **Form Data:**
  - `file`: The file to be uploaded (e.g., text or image)
- **Response:**

  ```json
  {
      "fileId": "unique-file-id"
  }
  ```

### 2. Create TodoItem

- **Endpoint:** `/todo`
- **Method:** `POST`
- **Body:**

  ```json
  {
      "description": "Your todo description",
      "dueDate": "2023-12-31T23:59:59Z",
      "fileId": "unique-file-id" // Optional
  }
  ```

- **Response:**

  ```json
  {
      "id": "uuid",
      "description": "Your todo description",
      "dueDate": "2023-12-31T23:59:59Z",
      "fileId": "unique-file-id",
      "createdAt": "2023-01-01T00:00:00Z",
      "updatedAt": "2023-01-01T00:00:00Z"
  }
  ```

## Running Tests

Execute unit tests using Make:

```bash
make test
```

## Running Benchmarks

Execute benchmarks using Make:

```bash
make benchmark
```

## Security Considerations

- **File Validation:** The application validates the file type based on extensions to prevent uploading unsupported or potentially harmful files.
- **UUIDs:** Utilizes UUIDs for unique identification of TodoItems and files to ensure security and uniqueness.
- **Environment Isolation:** Uses Docker to isolate services and dependencies, enhancing security and ease of setup.

## Notes

- **LocalStack:** The application uses LocalStack to mock AWS services like S3 and SQS. This allows you to develop and test without needing actual AWS credentials.
- **Dependency Injection:** The project employs dependency injection for services like S3 and SQS, making it easier to mock these services during testing.
- **Testing Libraries:** Utilizes `testify` for writing unit tests and mocking interfaces.

## Contributing

Feel free to open issues or submit pull requests for improvements and bug fixes.

## License

This project is licensed under the MIT License.