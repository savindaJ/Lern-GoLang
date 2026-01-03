Structure

go-backend/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── server/
│   │   └── http.go
│   ├── modules/
│   │   ├── user/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   ├── repository.go
│   │   │   ├── model.go
│   │   │   └── routes.go
│   ├── middleware/
│   │   └── auth.go
│   ├── database/
│   │   └── postgres.go
│   └── utils/
├── migrations/
├── pkg/
│   └── response/
├── .env
├── go.mod
├── go.sum
├── Dockerfile
└── README.md


Why this structure?

cmd/ → app entry points

internal/ → protected business logic

modules/ → feature-based architecture (clean & scalable)

pkg/ → reusable packages

migrations/ → DB migrations

Dependencies 

go get github.com/gin-gonic/gin
go get github.com/spf13/viper
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/golang-jwt/jwt/v5
go get go.uber.org/zap
