## Simple boilerplate m-banking rest-api with go

### `docker-compose up --build`
Run and build the app with docker-compose.\
Server started at [http://localhost:8080](http://localhost:8080).

### `install docker`
Install docker here [https://www.docker.com/get-started/](https://www.docker.com/get-started/)

**features:**
- login & register
- create customer & bank account numbers
- deposit
- cashout
- transfer balance
- check balance
- check mutations

**tech-stack used:**
- golang
- gorilla/mux
- gorm
- postgresql
- rabbitmq
- docker

## `Structure of the Project`
```md
└── m-banking
    ├── cmd
    │   └── main.go
    ├── internal
    │   ├── adapter
    │   │   ├── delivery
    │   │   │   └── http
    │   │   │       └── handlers.go
    │   │   ├── pubsub
    │   │   │   └── consumer.go
    │   │   └── repository
    │   │       └── repository.go
    │   └── core
    │       ├── domain
    │       │   ├── dto
    │       │   │   └── dto.go
    │       │   └── models
    │       │       └── models.go
    │       ├── ports
    │       │   └── ports.go
    │       └── usecase
    │           └── usecase.go
    ├── migrations
    │   └── migrations.go
    ├── pkg
    │   ├── bcrypt
    │   │   └── bcrypt.go
    │   ├── database
    │   │   └── sql
    │   │       └── postresql.go
    │   ├── jwt
    │   │   └── jwt.go
    │   └── rabbitmq
    │       └── rabbitmq.go
    ├── routes
    │   └── routes.go
    ├── .env
    ├── docker-compose.yml
    ├── Dockerfile
    ├── go.mod
    └── go.sum
```
