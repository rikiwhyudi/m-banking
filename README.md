## Simple boilerplate m-banking rest-api with go

### `docker-compose up --build`
Run and build the app with docker-compose.\
Server started at [http://localhost:8080](http://localhost:8080).

### `install docker`
Install docker here [https://www.docker.com/get-started/](https://www.docker.com/get-started/)

**feature:**
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
   ├── pkg
   ├── routes
   ├── docker-compose.yml
   ├── Dockerfile
   ├── go.mod
   ├── go.sum
   ├── .env
   └── internal
       ├── adapters
       │   │
       │   ├── delivery/http
       │   │   └── handler.go
       │   │
       │   └── repository
       │       └── repository.go
       │
       ├── core
       │   ├── domain/models
       │   │   └── model.go
       │   │   
       │   ├── ports
       │   │   └── ports.go
       │   │   
       │   └── usecase
       │       └── usecase.go
       └── dto
       │ 
       └── infrastructure
       │ 
       └── migrations
```

## Learn More
Other reference: [https://github.com/LordMoMA/Hexagonal-Architecture](https://github.com/LordMoMA/Hexagonal-Architecture)
