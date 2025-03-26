# 🛒 E-commerce Microservices with Golang, FastAPI, Flask, and Node.js

This project demonstrates a simple **microservices architecture** with **Golang, FastAPI, Flask, and Node.js**.  
Each service handles different aspects of an e-commerce system and communicates via REST APIs.  
The entire system runs with **Docker Compose**.

---

## 📌 Microservices Overview
| Service          | Technology | Port  | Description                     |
|-----------------|------------|------|--------------------------------|
| **API Gateway**  | Golang (Gin) | `5000` | Routes requests & aggregates data |
| **User Service** | Golang (Gin) | `5001` | Handles user details           |
| **Order Service** | FastAPI     | `5002` | Handles order details          |
| **Payment Service** | Flask    | `5003` | Handles payment details        |
| **Product Service** | Node.js (Express) | `5004` | Handles product details        |

---

## 🚀 How to Run
### **1️⃣ Clone the Repository**
```sh
git clone https://github.com/sajosam/ecommerce-microservices.git
cd ecommerce-microservices
```

### **2️⃣ Build & Start All Services**
```sh
docker-compose up --build
```

---

## 📌 API Endpoints
### 🔹 **User Service (Golang)**
- `GET /users/{id}` → Get user details  

### 🔹 **Order Service (FastAPI)**
- `GET /orders/{user_id}` → Get user orders  

### 🔹 **Payment Service (Flask)**
- `GET /payments/{user_id}` → Get user payments  

### 🔹 **Product Service (Node.js)**
- `GET /products/{user_id}` → Get user products  

### 🔹 **API Gateway (Golang)**
- `GET /all-details/{user_id}` → Fetch all details in a single response

---

## 🔥 Testing the APIs
### **Get User Details**
```sh
curl http://localhost:5000/users/1
```

### **Get User Orders**
```sh
curl http://localhost:5000/orders/1
```

### **Get User Payments**
```sh
curl http://localhost:5000/payments/1
```

### **Get User Products**
```sh
curl http://localhost:5000/products/1
```

### **Get All User Data in a Single JSON**
```sh
curl http://localhost:5000/all-details/1
```

---

## 📦 Docker Setup
Each microservice has a **Dockerfile** and is managed using **Docker Compose**.

To stop all services, run:
```sh
docker-compose down
```

---

## 📌 Tech Stack Used
✅ **Golang (Gin) - API Gateway & User Service**  
✅ **FastAPI - Order Service**  
✅ **Flask - Payment Service**  
✅ **Node.js (Express) - Product Service**  
✅ **Docker & Docker Compose**  

---


🚀 **Happy Coding!** 🔥

