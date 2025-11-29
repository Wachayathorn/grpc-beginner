# 🚀 gRPC Microservices Learning Project

โปรเจกต์เรียนรู้การสร้าง Microservices ด้วย Go และ gRPC สำหรับ beginners

---

## 📚 สิ่งที่จะได้เรียนรู้

1. ✅ **Protocol Buffers** - กำหนด API contract
2. ✅ **gRPC Server/Client** - สร้าง microservices ที่สื่อสารกัน
3. ✅ **Service Communication** - Users service เรียก Addresses service
4. ✅ **Dependency Injection** - แยก layer และ inject dependencies
5. ✅ **Clean Architecture** - handler → business → client

---

## 🏗️ Architecture

```
┌─────────────────┐         gRPC           ┌─────────────────┐
│  Users Service  │ ───────────────────→   │ Addresses       │
│  Port: 50051    │                        │ Service         │
│                 │                        │ Port: 50052     │
│  - ListUsers    │ ← gRPC Client calls    │ - ListAddresses │
└─────────────────┘                        └─────────────────┘
```

---

## 📁 Project Structure

```
grpc-beginner/
├── users/                    # Users microservice
│   ├── proto/
│   │   └── user.proto       # Proto definition
│   ├── pb/                  # Generated code (don't edit!)
│   ├── handler/             # gRPC handlers
│   │   └── handler.go
│   ├── business/            # Business logic
│   │   ├── business.go
│   │   └── model.go
│   ├── main.go              # Entry point
│   └── makefile
│
└── addresses/               # Addresses microservice
    ├── proto/
    │   └── address.proto
    ├── pb/
    ├── handler/
    ├── client/              # Client for other services to use
    └── main.go
```

---

## 🎯 Step-by-Step Guide

### **Step 1: ติดตั้ง Tools**

```bash
# 1. ติดตั้ง Protocol Buffers compiler
brew install protobuf

# 2. ติดตั้ง Go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 3. ตรวจสอบ
protoc --version
```

---

### **Step 2: สร้าง Proto File**

```protobuf
// users/proto/user.proto
syntax = "proto3";

package users.v1;
option go_package = "github.com/YOUR_USERNAME/grpc-beginner/users/pb/proto;pb";

service UserService {
  rpc ListUsers (ListUsersRequest) returns (ListUsersResponse);
}

message ListUsersRequest {
  int32 page = 1;
  int32 page_size = 2;
}

message ListUsersResponse {
  repeated User users = 1;
  int32 total = 2;
}

message User {
  int32 id = 1;
  string name = 2;
  string email = 3;
}
```

---

### **Step 3: Generate Go Code**

```bash
cd users

# Compile proto file
protoc --go_out=pb --go_opt=paths=source_relative \
       --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
       proto/user.proto

# หรือใช้ makefile
make proto
```

**ผลลัพธ์:**
- `pb/proto/user.pb.go` - Message definitions
- `pb/proto/user_grpc.pb.go` - gRPC service code

---

### **Step 4: สร้าง Handler**

```go
// users/handler/handler.go
package handler

import (
    "context"
    pb "github.com/YOUR_USERNAME/grpc-beginner/users/pb/proto"
)

type UserHandler struct {
    pb.UnimplementedUserServiceServer
}

func NewUserHandler() *UserHandler {
    return &UserHandler{}
}

func (h *UserHandler) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
    return &pb.ListUsersResponse{
        Users: []*pb.User{
            {Id: 1, Name: "John", Email: "john@example.com"},
        },
        Total: 1,
    }, nil
}
```

---

### **Step 5: สร้าง Server**

```go
// users/main.go
package main

import (
    "log"
    "net"
    
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    
    "github.com/YOUR_USERNAME/grpc-beginner/users/handler"
    pb "github.com/YOUR_USERNAME/grpc-beginner/users/pb/proto"
)

func main() {
    // 1. สร้าง handler
    userHandler := handler.NewUserHandler()
    
    // 2. สร้าง gRPC server
    grpcServer := grpc.NewServer()
    
    // 3. ลงทะเบียน service
    pb.RegisterUserServiceServer(grpcServer, userHandler)
    
    // 4. เปิด reflection (สำหรับ grpcurl)
    reflection.Register(grpcServer)
    
    // 5. Start server
    listener, _ := net.Listen("tcp", ":50051")
    log.Println("🚀 Server running on :50051")
    grpcServer.Serve(listener)
}
```

---

### **Step 6: รัน Server**

```bash
cd users
go run main.go
```

---

### **Step 7: ทดสอบด้วย grpcurl**

```bash
# ติดตั้ง grpcurl
brew install grpcurl

# ดู services ที่มี
grpcurl -plaintext localhost:50051 list

# เรียก ListUsers
grpcurl -plaintext -d '{
  "page": 1,
  "page_size": 10
}' localhost:50051 users.v1.UserService/ListUsers
```

---

## 🔗 Service-to-Service Communication

### **ใช้ Client ใน Users Service**

```go
// users/main.go
func main() {
    // สร้าง address client
    addressClient := addressclient.New()
    
    // Inject เข้า business logic
    userBusiness := business.New(addressClient)
    
    // Inject เข้า handler
    userHandler := handler.NewUserHandler(userBusiness)
    
    // สร้าง server...
}
```

---

## 🎓 Key Concepts

### **1. Protocol Buffers**
- **กำหนด contract** ระหว่าง services
- **Type-safe** - ตรวจสอบ type ตอน compile
- **Compact** - ขนาดเล็กกว่า JSON

### **2. gRPC**
- **Fast** - ใช้ HTTP/2
- **Bidirectional streaming** - ส่ง/รับพร้อมกันได้
- **Multiple languages** - Support หลายภาษา

### **3. Dependency Injection**
```go
addressClient := client.New()          // สร้าง dependency
business := business.New(addressClient) // Inject
handler := handler.New(business)        // Inject
```

### **4. Layers**
```
Handler   → รับ gRPC request/response
   ↓
Business  → Business logic, orchestration
   ↓
Client    → เรียก external services
```

---

## 🛠️ Makefile Commands

```bash
cd users

# Generate proto code
make proto

# Run server
make run
```

---

## 📊 Flow Diagram

```
1. Client
   ↓
2. Users Service (Handler)
   ↓
3. Business Logic
   ↓
4. Address Client → gRPC → Addresses Service (:50052)
   ↓
5. Combine data
   ↓
6. Return response
```

---

## ⚡ Quick Start

```bash
# 1. Clone
git clone https://github.com/Wachayathorn/grpc-beginner
cd grpc-beginner

# 2. Start Addresses service
cd addresses
go run main.go  # Port 50052

# 3. Start Users service (new terminal)
cd users
go run main.go  # Port 50051

# 4. Test
grpcurl -plaintext -d '{"page":1,"page_size":10}' \
  localhost:50051 users.v1.UserService/ListUsers
```

---

## 🐛 Common Issues

### **1. protoc not found**
```bash
brew install protobuf
```

### **2. Import path error**
```bash
go mod tidy
```

### **3. Cannot connect**
```bash
# ตรวจสอบว่า service รันอยู่
lsof -i :50051
```

---

## 📝 Files to Commit

### **✅ Should commit:**
- `proto/*.proto`
- `handler/`, `business/`, `main.go`
- `go.mod`, `go.sum`
- `README.md`, `Makefile`

### **❌ Don't commit:**
- `pb/` (generated code)
- Binary files

---

## 🎯 Next Steps

- [ ] เพิ่ม Database
- [ ] เพิ่ม Authentication
- [ ] เพิ่ม Error handling
- [ ] Add unit tests
- [ ] Dockerize

---

**Happy Learning! 🚀**
