.PHONY: run_inmemory_http
run_inmemory_http:
	memory_mode=inmemory transport_mode=http pg_pass=qwerty docker compose up --build

.PHONY: run_postgres_http
run_postgres_http:
	memory_mode=postgres transport_mode=http pg_pass=qwerty docker compose up --build

.PHONY: run_postgres_gRPC
run_postgres_gRPC:
	memory_mode=postgres transport_mode=gRPC pg_pass=qwerty docker compose up --build

.PHONY: run_inmemory_gRPC
run_inmemory_gRPC:
	memory_mode=inmemory transport_mode=gRPC pg_pass=qwerty docker compose up --build

.PHONY: run_default 
run_default:
	pg_pass=qwerty docker compose up --build


