## url_shortener

<br/>

## 🏷️ Description

All my projects is bricks 🧱 of road to becoming a developer ✨.

This project is dedicated to implementation url-shortener service. It's task for internship selection of ozon. We can see task-requirements is contains in TASK.md.

<p align="right"><a href="#url_shortener">Back to top ⬆️</a></p>

## 🎯 Solutions and Techniques

- Three-tier architecture: Transport -> Service -> Storage
- in-memory solutin using map + RWMutex 
- Auto application configuration using config parser
- Multi-level logging using zap logger
- Flexibility deploy with docker  
- Work with PostgreSQL usin pgx and squirrel libraries
- Unit-testing for business logic layer with testify;
- Testing gRPC server with evans tool 
- Simple run with makefile

<p align="right"><a href="#url_shortener">Back to top ⬆️</a></p>


## 🗂️ Table of Contents 
- [Description](#️-description)
- [Solutions and Techniques](#-solutions-and-techniques)
- [Table of Contents](#️-table-of-contents)
- [Working Tree](#-working-tree)
- [Getting Started](#️--getting-started)
- [API](#-api)
- [Usage](#-usage)
- [To do](#-to-do)
- [Contact](#-contact)

## 🌿 Working Tree
```
url_shortener
├── build
│   └── Dockerfile
├── cmd
│   └── url_shortener
│       └── main.go
├── config.yaml
├── docker-compose.yaml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── service
│   │   ├── encoder
│   │   │   ├── encoder.go
│   │   │   └── encoder_test.go
│   │   └── service.go
│   ├── storage
│   │   ├── inmemory
│   │   │   └── inmemory.go
│   │   └── postgres
│   │       └── postgres.go
│   └── transport
│       ├── gRPC
│       │   ├── gRPCHandler
│       │   │   └── gRPCHandler.go
│       │   └── gRPCServer
│       │       └── gRPCServer.go
│       └── http
│           ├── httpHandler
│           │   └── htppHandler.go
│           └── httpServer
│               └── httpServer.go
├── makefile
├── pkg
│   ├── logger
│   │   └── logger.go
│   └── validator
│       └── validator.go
├── proto
│   ├── url_shortener_grpc.pb.go
│   ├── url_shortener.pb.go
│   └── url_shortener.proto
├── README.md
└── TASK.md
```

<p align="right"><a href="#url_shortener">Back to top ⬆️</a></p>


## 🛠️  Getting Started

0. Install all required dependencies(Go, Docker and etc)

1. Clone the repository  

```bash
git clone https://github.com/VlPakhomov/url_shortener
```   

2. You can build and run containers with default settings and with the database using the following commands:
```
# create and compose up with default settings
 make run_default
```

3. You can choose memory_mode and transport_mode using another suffix run_{memory_mode}_{transport_mode} 

```
# create and compose up with memory_mode=inmemory and transport_mode=http 
 make run_inmemory_http
``` 

```
# create and compose up with memory_mode=postgres and transport_mode=gRPC 
 make run_postgres_gRPC
```

<p align="right"><a href="#url_shortener">Back to top ⬆️</a></p>

## ⛮ API

Application have two endpoints: 

- `POST /api/shorten-url/`
    - Handler for url convert to shorten url 
    - Request is plain text with UTF-8 that contains url
    - Response is also plain text with short url already

- `GET /api/get-url/`
  - Handler use to get original url 
  - Url passed by query string (example: api/get-url/?url=xaw1234gre)
  - Response is plain text with original url 

## 🧩 Usage

We can access the service with transport_mode=http using curl utility:
 - ```
    curl -X POST "http://localhost:8080/api/shorten-url/" -H "Content-Type: text/plain; charset=utf-8" -d "http://localhost:8080/api/" 
    ```
 - ```
    curl -X GET "http://localhost:8080/api/get-url/?url=____aRirKl" 
    ```

If you want access the service with transport_mode=gRPC,  use evans utility: 

- ```
  # connect to gRCP server on 8080 port
  evans proto/url_shortener.proto -p 8080 
  ``` 

- ```
  # call endpoint gRCP server 
  url_shortener.GrpcHandler@127.0.0.1:8080> call GetUrl
  ``` 

- ```
  # pass parameters for method 
  rawShortUrl (TYPE_STRING) => ____aowuMS
  ``` 

- ```
  # get a response 
  command call: rpc error: code = Unknown desc = url with ____aowuMS shortUrl doesn't exist | transportMode=gRPC
  ``` 

<p align="right"><a href="#url_shortener">Back to top ⬆️</a></p>

## 📌 To do 

- gRPC + evans ✅
- Makefile ✅
- Unit test for business logic layer ✅
- Testing transport and service by gomock
- Сleaning of unused Url
- .....

<p align="right"><a href="#url_shortener">Back to top ⬆️</a></p>



## 📫 Contact  

Vladislav Pakhomov - [@VlPakhomov](https://t.me/VlPakhomov) - [vladislavpakhomov03@gmail.com](mailto:vladislavpakhomov03@gmail.com)

Project Link: https://github.com/VlPakhomov/url_shortener

<p align="right"><a href="#url_shortener">Back to top ⬆️</a></p>


