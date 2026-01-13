```shell
go install github.com/swaggo/swag/cmd/swag@latest
cd backend/internal/controller
swag init -g index.go --pdl 1
```