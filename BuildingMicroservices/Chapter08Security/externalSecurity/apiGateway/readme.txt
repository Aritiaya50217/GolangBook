how to run 
1. รัน microservices
    go run ./user_service/user_service.go
    go run ./order_service/order_service.go
2. รัน API Gateway
    go run ./gateway/gateway.go
3. สร้าง token ใหม่
    go run ./generate_jwt/generate_jwt.go
4. ทดสอบด้วย curl
    curl -H "Authorization: Bearer <token>" http://localhost:8080/users/123



