cd product-api

go build

go run main.go

```
go-api
├─ config
│  ├─ config.go
│  └─ config.json
├─ crud
│  ├─ controller
│  │  └─ crud_controller.go
│  ├─ entity
│  │  ├─ crud_model.go
│  │  └─ env.go
│  ├─ repository
│  │  └─ crud_repository.go
│  ├─ repository.go
│  ├─ usecase
│  │  └─ crud_usecase.go
│  └─ usecase.go
├─ db
│  └─ mongo.go
├─ go.mod
├─ go.sum
├─ main.go
└─ router
   └─ router.go

```