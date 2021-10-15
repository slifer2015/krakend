### run krakend
docker run -p 7070:7070 -v $PWD:/etc/krakend/ devopsfaith/krakend run --config krakend.json

### run limit service
go run cmd/limit/main.go

### run auth service
go run cmd/auth/main.go

### run backend
go run cmd/backend/main.go

## test protected api
POST localhost:7070/api/v1/protected-api
Authorization header - > correct header
