# JWT Gin Authentication Service

A small Go REST API built with Gin that demonstrates user registration, login, and protected routes using JWT authentication. This repository uses MySQL via GORM for user storage and bcrypt to hash passwords.

## Features

- Register new users with username and password
- Login with username/password
- Generate JWT token on successful login
- Protect API routes using JWT middleware
- Retrieve current authenticated user data

## Project Structure

- `main.go` - application entry point and route setup
- `controllers/auth.go` - register, login, and current user handlers
- `models/setup.go` - database connection and migration setup
- `models/user.go` - user model and authentication logic
- `utils/token/token.go` - JWT generation and validation helpers
- `middlewares/middlewares.go` - JWT middleware for protected routes

## Requirements

- Go 1.20+ (or compatible)
- MySQL database
- `go` command available in your PATH

## Environment Variables

Create a `.env` file in the project root with the following values:

```env
DB_DRIVER=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=jwt_gin
TOKEN_HOUR_LIFESPAN=24
API_SECRET=your_secret_key
```

Adjust the values as needed for your environment.

## Install Dependencies

```bash
go mod download
```

## Run the Application

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

## API Endpoints

### Register

- Method: `POST`
- URL: `/api/register`
- Body:
  ```json
  {
    "username": "example",
    "password": "password123"
  }
  ```
- Response: registration success message

### Login

- Method: `POST`
- URL: `/api/login`
- Body:
  ```json
  {
    "username": "example",
    "password": "password123"
  }
  ```
- Response: JWT token

### Get Current User

- Method: `GET`
- URL: `/api/admin/user`
- Header: `Authorization: Bearer <token>`
- Response: authenticated user data

## Notes

- Passwords are hashed before saving to the database using bcrypt.
- JWT tokens are signed with `API_SECRET` and expire after the time specified by `TOKEN_HOUR_LIFESPAN`.

## License

This repository is provided as-is for learning and demonstration purposes.
