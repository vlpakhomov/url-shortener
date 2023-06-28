.PHONY: build
build:
	docker compose build

.PHONY: build_inmemory_http
build_inmemory_http:
	memory_mode=inmemory transport_mode=http pg_pass=qwerty docker compose up --build

.PHONY: build_postgres_http
build_postgres_http:
	memory_mode=postgres transport_mode=http pg_pass=qwerty docker compose up --build

.PHONY: build_postgres_gRPC
build_postgres_gRPC:
	memory_mode=postgres transport_mode=gRPC pg_pass=qwerty docker compose up --build

.PHONY: build_inmemory_gRPC
build_inmemory_gRPC:
	memory_mode=inmemory transport_mode=http pg_pass=qwerty docker compose up --build

.PHONY: cold_compose_up 
cold_run:
	docker system prune && pg_pass=qwerty docker compose up --build

.PHONY: compose_up
compose_up:
	pg_pass=qwerty docker compose up


