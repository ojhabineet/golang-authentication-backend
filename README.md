# Golang Authentication API

A backend authentication system built using Go and Gin framework.  
This project implements user signup, login, JWT-based authentication, password hashing, and role-based access control without using a database.

## Features

- User Signup API
- User Login API
- Password hashing using bcrypt
- JWT authentication
- Protected routes
- Role-based authorization (User/Admin)
- User profile access
- Admin-only access to all users
- In-memory data storage

## Tech Stack

- Go (Golang)
- Gin Web Framework
- JWT (JSON Web Tokens)
- bcrypt for password hashing

## Project Structure

```bash
backend-auth/
│
├── main.go
├── go.mod
│
├── handlers/
│   └── auth.go
│
├── middleware/
│   └── auth.go
│
├── models/
│   └── user.go
│
├── storage/
│   └── storage.go
│
└── utils/
    └── jwt.go
```

## Installation and Setup

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/golang-auth-api.git
cd golang-auth-api
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the server

```bash
go run main.go
```

Server will start at:

```bash
http://localhost:8080
```

## API Endpoints

### 1. Signup

**Endpoint**

```http
POST /signup
```

**Request Body**

```json
{
  "name": "Bineet",
  "email": "bineet@example.com",
  "password": "123456"
}
```

**Response**

```json
{
  "message": "Signup successful"
}
```

---

### 2. Login

**Endpoint**

```http
POST /login
```

**Request Body**

```json
{
  "email": "bineet@example.com",
  "password": "123456"
}
```

**Response**

```json
{
  "token": "your_jwt_token"
}
```

---

### 3. Get Profile

**Endpoint**

```http
GET /profile
```

**Headers**

```txt
Authorization: Bearer your_jwt_token
```

**Response**

```json
{
  "id": 1,
  "name": "Bineet",
  "email": "bineet@example.com",
  "role": "user"
}
```

---

### 4. Get All Users (Admin Only)

**Endpoint**

```http
GET /users
```

**Headers**

```txt
Authorization: Bearer your_jwt_token
```

## Authentication Flow

### Signup Flow

```txt
Client → Signup Request → Validate Input → Hash Password → Store User
```

### Login Flow

```txt
Client → Verify Credentials → Generate JWT → Return Token
```

### Protected Route Flow

```txt
Client → Send JWT → Middleware Verifies Token → Access Granted
```

## Notes

- User data is stored in memory and will be lost when the server restarts.
- This project is intended for learning and demonstration purposes.
- Database integration (such as PostgreSQL) can be added later for persistence.

## Future Improvements

- Database integration with PostgreSQL
- Refresh token support
- Password reset functionality
- Email verification
- Docker support
- Logging and monitoring
- Rate limiting
- API documentation with Swagger

## Author

Bineet Ojha
