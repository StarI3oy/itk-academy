Prerequisites (recommended):

```
go 1.25 >=
make v4.4.1 >=
docker 28.1.1 >=
k6 v1.7.1>=
golang-migrate v4.19.1 >=
```


Перед началом работы:

1) Запустить ```make compose```
2) Запустить ```make migrate```

Для запуска тестов:
1) Запустить ```make test```
2) Для нагрузки ```make test-k6``` (предварительно нужно установить k6)

Структура проекта:

```
build/
│   └── dockerfile*
├── cmd/
│   └── app/
│       └── main.go*
├── internal/
│   ├── app/
│   │   └── server.go*
│   ├── config/
│   │   └── config.go*
│   ├── deps/
│   │   └── deps.go*
│   ├── domain/
│   │   ├── entity/
│   │   │   └── wallet.go*
│   │   ├── ports/
│   │   │   └── walletRepo.go*
│   │   └── service/
│   │       └── wallet_service.go*
│   └── transport/
│       └── http/
│           ├── dto/
│           │   └── wallet.go*
│           └── wallet_handler.go*
├── migrations/
│   ├── 000000_create_all_tables.down.sql*
│   ├── 000000_create_all_tables.up.sql*
│   ├── 000001_seed.down.sql*
│   └── 000001_seed.up.sql*
├── pkg/
│   ├── db/
│   │   └── postgres.go*
│   ├── errors/
│   │   └── wallet.go*
│   ├── logger/
│   │   └── logger.go*
│   ├── model/
│   │   └── wallet/
│   │       └── wallet.go*
│   └── repository/
│       └── wallet/
│           ├── wallet.go*
│           └── wallet_operations.go*
├── tests/
│   ├── concurrency/
│   │   └── concurrency_test.go*
│   ├── helpers/
│   │   └── util.go*
│   ├── http_t/
│   │   ├── http_test.go*
│   │   └── stub.go*
│   ├── integration/
│   │   ├── wallet_repo_test.go*
│   │   └── wallet_service_test.go*
│   ├── load/
│   │   └── wallet-test.js*
│   └── testdb.go*
```