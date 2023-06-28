## url_shortener

<br/>

## 🏷️ Description

All my projects is bricks 🧱 of road to becoming a developer ✨.

This project is dedicated to implementation url-shortener service. It's task for internship selection. We can see task-requirements is contains in TASK.md ⌛.

<p align="right"><a href="#url_shortener">Back to top ⬆️</a></p>

## 🎯 Solutions and Techniques

- Three-tier architecture: Transport -> Service -> Storage
- in-memory solutin using map + RWMutex 
- Auto application configuration using config parser
- Multi-level logging using zap logger
- Flexibility deploy with Docker 
- Work with PostgreSQL usin pgx and squirrel libraries

<p align="right"><a href="#url_shortener">Back to top ⬆️</a></p>


## 🗂️ Table of Contents 
- [url\_shortener](#url_shortener)
- [🏷️ Description](#️-description)
- [🎯 Solutions and Techniques](#-solutions-and-techniques)
- [🗂️ Table of Contents](#️-table-of-contents)
- [🌿 Working Tree](#-working-tree)
- [🛠️  Getting Started](#️--getting-started)
- [⛮ API](#-api)
- [🧩 Usage](#-usage)
- [📌 To do](#-to-do)
- [📫 Contact](#-contact)

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
│   │   │   └── encoder.go
│   │   └── service.go
│   ├── storage
│   │   ├── inmemory
│   │   │   └── inmemory.go
│   │   └── postgres
│   │       └── postgres.go
│   └── transport
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

2. Set configuration in config.yaml(MEMORY_MODE, TRANSPORT_MODE and etc)
   


3. Docker build
```
 pg_pass=qwerty docker compose up {--build| if it's first start} {-d| disable logs}
```

4. Set logs
   
```
 docker logs -f url_shortener
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

## 🧩 Usage

We can access the service using curl utility:
 - ```
    curl -X POST "http://localhost:8080/api/shorten-url/" -H "Content-Type: text/plain; charset=utf-8" -d "http://localhost:8080/api/" 
    ```
 - ```
    curl -X GET "http://localhost:8080/api/get-url/?url=____aRirKl" 
    ```

<p align="right"><a href="#url_shortener">Back to top ⬆️</a></p>

## 📌 To do 

- gRPC ✅
- Makefile 
- Unit and E2E test 
- Сleaning of unused Url
- .....

<p align="right"><a href="#url_shortener">Back to top ⬆️</a></p>



## 📫 Contact  

Vladislav Pakhomov - [@VlPakhomov](https://t.me/VlPakhomov) - [vladislavpakhomov03@gmail.com](mailto:vladislavpakhomov03@gmail.com)

Project Link: https://github.com/VlPakhomov/url_shortener

<p align="right"><a href="#url_shortener">Back to top ⬆️</a></p>


