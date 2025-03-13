# Ecommerce API

## Description
Simple API for technical test PT Praweda Sarana Informatika with discount voucher system and earning points calculation

## Technology Used
- Go (Golang)
- Gin Framework

## How To Use
1. Make sure Go is installed on your system
2. Clone this repository:
   ```sh
   git clone https://github.com/adicipta/TechnicalTestPSI/tree/7a1f2175bfb3c7c6165fa3a69e5b9b8d89f46abc/soal1
   ```
3. Go to project directory:
   ```sh
   cd ecommerce-checkout
   ```
4. Run the application:
   ```sh
   go run main.go
   ```

## Endpoint API
### 1. Checkout
**URL:** `POST /checkout`

**Request Body With Voucher:**
```json
{
  "user_id": 1,
  "item_id": 1,
  "voucher_code": "Discount50"
}
```

**Response (200 - Success with voucher):**
```json
{
  "message": "Checkout Success",
  "total_price": 2500000,
  "earned_point": 50000
}
```

**Request Body Without Voucher:**
```json
{
  "user_id": 1,
  "item_id": 1,
  "voucher_code": ""
}
```

**Response (200 - Success without voucher):**
```json
{
  "message": "Checkout Success",
  "total_price": 5000000,
  "earned_point": 0
}
```

**Response (404 - User not found):**
```json
{
  "error": "User not found"
}
```

**Response (404 - Item not found):**
```json
{
  "error": "Item not found"
}
```

**Response (404 - Voucher not found):**
```json
{
  "error": "Voucher not found"
}
```