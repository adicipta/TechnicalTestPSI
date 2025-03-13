# Auth API

## Description
Simple API for technical test PT Praweda Sarana Informatika with cookies and jwt auth

## Technology Used
- Golang (Go)
- Gin Framework
- JWT-Go v5
- UUID

## How To Use
1. Make sure Go is installed on your system
2. Clone this repository:
   ```sh
   git clone https://github.com/adicipta/TechnicalTestPSI/tree/7d1c0f880b8b8d4ef8a317d2e5196fb7771ede36/soal2
   ```
3. Go to project directory:
   ```sh
   cd soal2
   ```
4. Copy .env-example .env
5. Create .env file and set the following variables:
   ```
   JWT_SECRET=YOUR_SECRET_KEY
   ```
6. Run the application:
   ```sh
   go run main.go
   ```

## Endpoint API
### 1. Login
**URL:** `POST /login`

**Request Body With Username:**
```json
{
  "username": "your_username",
}
```

**Response (200 - Success with voucher):**
```json
{
  "message": "Login successful",
  "session_exp": 17709012,
  "token": "exampletoken"
}
```

**URL** `GET /protected`

**Headers:**
```
Cookie: token=your_jwt_token
```

**Response (200 - Success get protected):**
```json
{
  "id": "uuid-v4-id",
  "username": "your_username",
  "session_exp": 17709012
}
```

**Response (400 - Username is required):**
```json
{
  "error": "Username is required"
}
```

**Response (401 - Token not found):**
```json
{
  "error": "Token not found"
}
```