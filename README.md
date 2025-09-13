# 🔑 Locksmith (Crypto Key Management Service)

Locksmith is a **crypto key management & distribution service** based on gRPC, Redis, dan RabbitMQ. This project initially started from BTC key generator, then developed into modern service for **crypto enthusiasts dan developer**.

## Running

```bash
go run main.go btc --bip44 -p=jerapahimut -i=10
```

## Running with gRPC save mode

save the generated keys, mnemonic, and wallet address into database bt gRPC server

```bash
go run .\main.go eth -p=hesoyam --save-remote
```

## Start gRPC Server

```bash
go run cmd/server/main.go
```

## Migrations

Download and install `migrate` binary to create and run migrations

```powershell
winget install golang-migrate.migrate
```

If winget is not available, download from:
[https://github.com/golang-migrate/migrate/releases](https://github.com/golang-migrate/migrate/releases)
→ Download migrate.windows-amd64.zip, extract, and add to PATH.

Then run the command bellow to generate migration files

```bash
migrate create -ext sql -dir internal/core/migrations -seq create_wallets_table
```

Then run the following command to create new table

```bash
migrate -database "mysql://youruser:yourpassword@tcp(localhost:3306)/dbname" -path internal/core/migrations up
```

delete table

```bash
migrate -database "mysql://youruser:yourpass@tcp(localhost:3306)/dbname" -path internal/core/migrations down
```

## 🧑‍💻 Author

Built with ❤️ by Wahyu Krisna Aji
