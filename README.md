# E-Wallet RESTful API

A production-ready e-wallet system built with Go, featuring user management, wallet operations, and secure money transfers with ACID compliance and race condition handling.

[![Go Version](https://img.shields.io/badge/Go-1.24-blue.svg)](https://golang.org)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](https://www.docker.com/)

## 🚀 Quick Start

```bash
# Clone and setup
git clone <repository-url>
cd mywallet
cp .env.example .env

# Run with Docker (easiest)
docker-compose up -d

# API is ready at http://localhost:8080
```

Test it:
```bash
# Register a user
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com","password":"password123"}'

# Login and get token
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"alice@example.com","password":"password123"}'
```

## 📚 Table of Contents
- [Architecture](#-architecture)
- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Getting Started](#-getting-started)
- [API Endpoints](#-api-endpoints)
- [Database Schema](#-database-schema)
- [Migration Management](#-migration-management)

## 🛠️ Tech Stack

**Backend Framework:**
- Go 1.24
- Gin Web Framework v1.11.0 (HTTP router)

**Database:**
- MySQL 8.0
- GORM v1.31.1 (ORM)

**Authentication & Security:**
- JWT (golang-jwt/jwt v5.3.1)
- Bcrypt (golang.org/x/crypto)
- go-playground/validator v10.27.0

**Configuration & Environment:**
- Viper v1.21.0 (config management)

**DevOps:**
- Docker & Docker Compose
- golang-migrate (database migrations)

**Architecture Pattern:**
- Clean Architecture (layered)
- Repository Pattern
- Dependency Injection

## 🏗️ Architecture

Clean layered architecture with clear separation of concerns:

```
mywallet/
├── controller/       # HTTP handlers (presentation layer)
├── usecase/          # Business logic & orchestration
├── repository/       # Data access layer
├── model/            # Database models
├── dto/              # Request/Response DTOs
├── middleware/       # Auth, CORS, error handling
├── apperror/         # Application-specific errors
├── constant/         # Application constants
├── utils/            # App utilities (converter, pagination)
├── pkg/              # Reusable packages (jwt, hash, validator)
├── config/           # Configuration management
├── server/           # Server initialization
└── migrations/       # Database migrations
```

**Architecture Pattern**: Controller → Usecase → Repository

## ✨ Features

### 1. User Management
- ✅ User registration with email validation
- ✅ Secure login with JWT authentication
- ✅ Password hashing with bcrypt (cost=12)
- ✅ User profile retrieval

### 2. Wallet Management
- ✅ Automatic wallet creation on user registration
- ✅ Balance inquiry
- ✅ Top-up functionality with validation
- ✅ Decimal precision for financial data (19,2)

### 3. Transaction Management
- ✅ Transfer money between users
- ✅ Transaction history with pagination
- ✅ ACID compliance via database transactions
- ✅ Race condition prevention (SELECT FOR UPDATE)
- ✅ Transaction status tracking (PENDING/SUCCESS/FAILED)

### 4. Security Features (OWASP Compliant)
- ✅ JWT-based authentication with configurable expiration
- ✅ Password hashing with bcrypt
- ✅ SQL injection prevention (prepared statements via GORM)
- ✅ Input validation & sanitization
- ✅ Soft delete for data integrity
- ✅ UTC timestamps for consistency
- ✅ CORS middleware
- ✅ Centralized error handling

## 🚀 Getting Started

### Prerequisites

- **Docker & Docker Compose** (recommended) OR
- Go 1.24+
- MySQL 8.0+
- golang-migrate (for local development)

### Option 1: Docker (Recommended)

1. **Clone the repository**
```bash
git clone <repository-url>
cd mywallet
```

2. **Setup environment variables**
```bash
cp .env.example .env
# Edit .env and change JWT_SECRET to a secure random string
```

3. **Run with Docker Compose**
```bash
# Start all services (MySQL + migrations + app)
docker-compose up -d

# View logs
docker-compose logs -f app

# Stop services
docker-compose down

# Stop and remove volumes (⚠️ deletes database)
docker-compose down -v
```

The API will be available at `http://localhost:8080`

### Option 2: Local Development

1. **Clone the repository**
```bash
git clone <repository-url>
cd mywallet
```

2. **Install dependencies**
```bash
go mod download
```

3. **Setup environment variables**
```bash
cp .env.example .env
# Edit .env with your local MySQL configuration
```

4. **Install golang-migrate**
```bash
# Windows (Scoop)
scoop install migrate

# Windows (Chocolatey)
choco install golang-migrate

# MacOS
brew install golang-migrate

# Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/
```

5. **Create database**
```sql
CREATE DATABASE mywallet_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

6. **Run migrations**
```bash
migrate -path migrations -database "mysql://root:password@tcp(localhost:3306)/mywallet_db" up
```

7. **Run the application**
```bash
go run main.go
```

The server will start at `http://localhost:8080`

## � Makefile Commands

For convenience, common operations are available via Makefile:

```bash
# View all available commands
make help

# Application
make build          # Build binary
make run            # Run application
make test           # Run tests
make deps           # Download dependencies
make clean          # Clean build artifacts

# Docker
make docker-up      # Start all services
make docker-down    # Stop all services
make docker-clean   # Stop all services and remove volumes (deletes data)
make docker-logs    # View logs
make docker-restart # Restart services

# Database
make migrate-up     # Apply migrations
make migrate-down   # Rollback migration
make migrate-create NAME=table_name  # Create new migration
```

## �📡 API Endpoints

### Base URL
- Local: `http://localhost:8080`
- Production: `https://api.yourdomain.com`

### Response Format
All responses follow this structure:
```json
{
  "status": "success|error",
  "data": { ... },      // On success
  "meta": { ... },      // Optional: pagination, rate limits, etc.
  "error": "message"    // On error
}
```

### Authentication (Public)

#### Register User
```http
POST /api/auth/register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "SecurePass123"
}

Response (200 OK):
{
  "status": "success",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "j***@example.com",
    "created_at": "2026-02-12T10:00:00Z"
  }
}
```

#### Login
```http
POST /api/auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "SecurePass123"
}

Response (200 OK):
{
  "status": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "j***@example.com",
      "created_at": "2026-02-12T10:00:00Z"
    }
  }
}
```

### User (Protected - Requires JWT)

#### Get Profile
```http
GET /api/users/profile
Authorization: Bearer <your-jwt-token>

Response (200 OK):
{
  "status": "success",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "j***@example.com",
    "created_at": "2026-02-12T10:00:00Z"
  }
}
```

### Wallet (Protected - Requires JWT)

#### Get Balance
```http
GET /api/wallets/balance
Authorization: Bearer <your-jwt-token>

Response (200 OK):
{
  "status": "success",
  "data": {
    "id": 1,
    "user_id": 1,
    "balance": 1000000.00,
    "created_at": "2026-02-12T10:00:00Z",
    "updated_at": "2026-02-12T15:30:00Z"
  }
}
```

#### Top Up Wallet
```http
POST /api/wallets/topup
Authorization: Bearer <your-jwt-token>
Content-Type: application/json

{
  "amount": 500000.00
}

Response (200 OK):
{
  "status": "success",
  "data": {
    "wallet_id": 1,
    "new_balance": 1500000.00,
    "transaction_id": 42
  }
}
```

### Transactions (Protected - Requires JWT)

#### Transfer Money
```http
POST /api/transactions/transfer
Authorization: Bearer <your-jwt-token>
Content-Type: application/json

{
  "receiver_email": "jane@example.com",
  "amount": 150000.00,
  "description": "Payment for services"
}

Response (200 OK):
{
  "status": "success",
  "data": {
    "transaction_id": 43,
    "sender_wallet_id": 1,
    "receiver_wallet_id": 2,
    "amount": 150000.00,
    "new_balance": 1350000.00,
    "status": "SUCCESS"
  }
}
```

#### Get Transaction History
```http
GET /api/transactions/history?page=1&limit=10
Authorization: Bearer <your-jwt-token>

Response (200 OK):
{
  "status": "success",
  "data": [
    {
      "id": 43,
      "transaction_type": "TRANSFER",
      "sender_wallet_id": 1,
      "receiver_wallet_id": 2,
      "amount": 150000.00,
      "status": "SUCCESS",
      "description": "Payment for services",
      "created_at": "2026-02-12T15:30:00Z"
    }
  ],
  "meta": {
    "page": 1,
    "limit": 10,
    "total": 25,
    "total_pages": 3
  }
}
```

### Error Responses

**Validation Error (400):**
```json
{
  "status": "error",
  "error": "Amount must be greater than zero"
}
```

**Unauthorized (401):**
```json
{
  "status": "error",
  "error": "Unauthorized access"
}
```

**Not Found (404):**
```json
{
  "status": "error",
  "error": "Wallet not found"
}
```

**Conflict (409):**
```json
{
  "status": "error",
  "error": "Insufficient balance for this transaction"
}
```

## 🗄️ Database Schema

### Users Table
- Primary Key: `id`
- Unique: `email`
- Fields: `name`, `password_hash`
- Timestamps: `created_at`, `updated_at`, `deleted_at` (soft delete)

### Wallets Table
- Primary Key: `id`
- Foreign Key: `user_id` → `users(id)` (UNIQUE)
- Fields: `balance` (DECIMAL 19,2)
- Constraint: `balance >= 0`
- Timestamps: `created_at`, `updated_at`, `deleted_at`

### Transactions Table
- Primary Key: `id`
- Foreign Keys: `sender_wallet_id`, `receiver_wallet_id` → `wallets(id)`
- Fields: `transaction_type` (TOPUP/TRANSFER), `amount`, `status` (PENDING/SUCCESS/FAILED), `description`
- Indexes: `created_at`, `sender_wallet_id`, `receiver_wallet_id`, `status`
- Timestamps: `created_at`, `updated_at`, `deleted_at`
- Note: All timestamps stored in UTC

## 🧪 Testing

### Postman Collection

Import `E-Wallet-API.postman_collection.json` into Postman for easy API testing. The collection includes:
- Auto-saves JWT token after login
- All endpoints with example payloads
- Environment variables for base URL

### Manual Testing with cURL

**Register:**
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com","password":"password123"}'
```

**Login:**
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"alice@example.com","password":"password123"}'
```

**Get Balance:**
```bash
curl -X GET http://localhost:8080/api/wallets/balance \
  -H "Authorization: Bearer <your-token>"
```

**Top Up:**
```bash
curl -X POST http://localhost:8080/api/wallets/topup \
  -H "Authorization: Bearer <your-token>" \
  -H "Content-Type: application/json" \
  -d '{"amount":1000000}'
```

**Transfer:**
```bash
curl -X POST http://localhost:8080/api/transactions/transfer \
  -H "Authorization: Bearer <your-token>" \
  -H "Content-Type: application/json" \
  -d '{"receiver_email":"bob@example.com","amount":50000,"description":"Payment"}'
```

## �📝 Migration Management

### Create New Migration
```bash
migrate create -ext sql -dir migrations -seq add_new_table

or

make migrate-create NAME=table_name
```

### Apply Migrations
```bash
migrate -path migrations -database "mysql://root:password@tcp(localhost:3306)/mywallet_db" up

or

make migrate-up
```

### Rollback Last Migration
```bash
migrate -path migrations -database "mysql://root:password@tcp(localhost:3306)/mywallet_db" down 1

or

make migrate-down
```

### Check Migration Version
```bash
migrate -path migrations -database "mysql://root:password@tcp(localhost:3306)/mywallet_db" version
```