test:
	make migrate-test
	go test ./tests/... -p 1

test-k6:
	k6 run tests/load/wallet-test.js

compose:
	docker-compose up -d


migrate:
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:6432/wallet?sslmode=disable" up

migrate-test: 
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5434/wallet_test?sslmode=disable" up