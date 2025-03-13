# Fetch API

## Description
Simple API for technical test PT Praweda Sarana Informatika for fetch API

## Technology Used
- Go (Golang)
- Gin Framework

## How To Use
1. Make sure Go is installed on your system
2. Clone this repository:
   ```sh
   git clone https://github.com/adicipta/TechnicalTestPSI/
   ```
3. Go to project directory:
   ```sh
   cd soal4
   ```
4. Run the application:
   ```sh
   go run main.go
   ```

## Endpoint API
### 1. User
**URL:** `GET /user`

**Response (200 - Success with user):**
```json
{
  "name": "Mr, Mads Hilleren",
  "location": "3355, Lille Bislett, Østenstad, Trøndelag, Norway",
  "email": "mads.hilleren@example.com",
  "age": 70,
  "phone": "81587974",
  "cell": "97044286",
  "picture": [
        "https://randomuser.me/api/portraits/men/81.jpg",
        "https://randomuser.me/api/portraits/med/men/81.jpg",
        "https://randomuser.me/api/portraits/thumb/men/81.jpg"
    ]
}
```