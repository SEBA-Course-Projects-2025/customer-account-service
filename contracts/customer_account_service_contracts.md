# Customer Account Service API Contracts

---

__Swagger API Documentation: https://customer-account-service.onrender.com/swagger/index.html__

## 0. Customer Login

__POST ```/auth/login```__

__Body:__

```json
{
  "email": "string",
  "password": "string"
}
```

__Response:__

```json
{
  "token": "jwt-token"
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request```
- ```401 Unauthorized```
- ```500 Internal Server Error```

## 1. Get Customer Account Information

__GET ```api/account```__

__Response:__

```json
{
  "id": "uuid",
  "email": "string",
  "name": "string",
  "phone": "string",
  "shippingAddress": "string"
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request```
- ```404 Not Found```
- ```500 Internal Server Error```

---

## 2. Update Customer Account Information

__PUT ```api/account```__

__Body:__

```json
{
  "email": "string",
  "name": "string",
  "phone": "string",
  "shippingAddress": "string"
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request```
- ```404 Not Found```
- ```500 Internal Server Error```

---

## 3. Modify Customer Account Information

__PATCH ```api/account```__

__Body:__

```json
{
  "email": "string", (optional)
  "name": "string", (optional)
  "phone": "string", (optional)
  "shippingAddress": "string" (optional)
}
```

__Response:__

```json
{
  "id": "uuid",
  "email": "string",
  "name": "string",
  "phone": "string",
  "shippingAddress": "string"
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request```
- ```404 Not Found```
- ```500 Internal Server Error```

---


## 4. Check Customer Orders

__GET ```api/orders```__

__Response:__

```json
{
  "order_id": "uuid",
  "customer_id": "uuid",
  "items": ["string"],
  "total_price": "float64",
  "status": "string",
  "date": "date"
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request```
- ```500 Internal Server Error```

---

## 5. Get One Order

__GET ```api/orders/:orderId```__

__Response:__

```json
{
  "orderId": "uuid",
  "customerId": "uuid",
  "items": [
    {
      "productId": "uuid",
      "product_name": "string",
      "quantity": "int",
      "image_url": "string",
      "unit_price": "float64"
    }
  ],
  "totalPrice": "float64",
  "status": "string",
  "date": "date"
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request```
- ```404 Not Found```
- ```500 Internal Server Error```

---

## 6. Cancel Orders

__PUT ```api/orders/:orderId```__

__Body:__

```json
{
  "status": "string"
}
```

__Status codes:__
- ```200 OK (success)```
- ```400 Bad Request```
- ```404 Not Found```
- ```500 Internal Server Error```

---










