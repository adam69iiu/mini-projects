run:
	go run ./cmd/server

migrate-up:
	goose -dir db/migrations postgres "host=localhost port=5432 user=adam password=adam1234 dbname=auth_db sslmode=disable" up

migrate-down:
	goose -dir db/migrations postgres "host=localhost port=5432 user=adam password=adam1234 dbname=auth_db sslmode=disable" down

migrate-status:
	goose -dir db/migrations postgres "host=localhost port=5432 user=adam password=adam1234 dbname=auth_db sslmode=disable" status

