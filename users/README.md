# Users Service 👥

Users microservice ที่ใช้ gRPC สำหรับจัดการข้อมูลผู้ใช้

## 📂 Project Structure

```
users/
├── service/
│   └── user.proto              # Protocol Buffer definitions
├── pb/                         # Generated protobuf code (auto-generated)
├── internal/
│   ├── domain/                 # Business entities
│   │   └── user.go
│   ├── repository/             # Data access layer
│   │   ├── user_repository.go
│   │   └── memory_repository.go
│   ├── usecase/                # Business logic
│   │   └── user_usecase.go
│   └── delivery/grpc/          # gRPC handlers
│       └── handler.go
├── cmd/server/
│   └── main.go                 # Application entry point
├── pkg/                        # Shared utilities
├── Makefile
└── README.md
```

## 🚀 Quick Start

### 1. Generate Protobuf Code
```bash
make proto
```

### 2. Run Server
```bash
make run
```

Server จะรันที่ `localhost:50051`

## 📝 Available Commands

| Command | Description |
|---------|-------------|
| `make proto` | Generate Go code จาก proto files |
| `make run` | รัน gRPC server |
| `make clean` | ลบ generated files |
| `make deps` | ติดตั้ง Go dependencies |
| `make test` | รัน tests |

## 🧪 Testing with grpcurl

### ติดตั้ง grpcurl
```bash
brew install grpcurl
```

### 1. List Services
```bash
grpcurl -plaintext localhost:50051 list
```

### 2. Create User
```bash
grpcurl -plaintext -d '{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "0812345678"
}' localhost:50051 users.v1.UserService/CreateUser
```

### 3. Get User
```bash
grpcurl -plaintext -d '{"id": 1}' localhost:50051 users.v1.UserService/GetUser
```

### 4. Update User
```bash
grpcurl -plaintext -d '{
  "id": 1,
  "name": "John Updated",
  "email": "john.updated@example.com",
  "phone": "0898765432"
}' localhost:50051 users.v1.UserService/UpdateUser
```

### 5. List Users
```bash
grpcurl -plaintext -d '{
  "page": 1,
  "page_size": 10
}' localhost:50051 users.v1.UserService/ListUsers
```

### 6. Delete User
```bash
grpcurl -plaintext -d '{"id": 1}' localhost:50051 users.v1.UserService/DeleteUser
```

## 🏗️ Architecture

Service นี้ใช้ **Clean Architecture** แบ่งเป็น 4 layers:

1. **Domain Layer** - Business entities (User)
2. **Repository Layer** - Data access (In-memory storage)
3. **Usecase Layer** - Business logic & validation
4. **Delivery Layer** - gRPC handlers

### Dependency Flow
```
Delivery → Usecase → Repository → Domain
```

## 🔧 Features

- ✅ CRUD operations (Create, Read, Update, Delete)
- ✅ Pagination support
- ✅ Email validation
- ✅ In-memory storage (ง่ายต่อการเปลี่ยนเป็น Database)
- ✅ Clean Architecture
- ✅ gRPC Reflection enabled

## 📦 Dependencies

- `google.golang.org/grpc` - gRPC framework
- `google.golang.org/protobuf` - Protocol Buffers

## 🔮 Next Steps

- [ ] เพิ่ม PostgreSQL database
- [ ] เพิ่ม authentication/authorization
- [ ] เพิ่ม logging middleware
- [ ] เพิ่ม error handling middleware
- [ ] เพิ่ม unit tests
- [ ] เพิ่ม Docker support
