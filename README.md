# 🔑 Locksmith (Crypto Key Management Service)

Locksmith adalah **crypto key management & distribution service** berbasis gRPC, Redis, dan RabbitMQ. Proyek ini awalnya dimulai dari simple BTC key generator, lalu dikembangkan jadi service modern untuk **crypto enthusiasts dan developer**.

---

## ✨ Features (MVP)

- Generate BTC key pair (private/public key).
- Simpan key sementara di Redis dengan TTL.
- Publish event ke RabbitMQ setiap kali key dibuat.
- Consumer sederhana untuk mendengarkan event `NewKeyGenerated`.
- Docker Compose untuk spin-up `app + redis + rabbitmq`.

---

## 🚀 Planned Features (Roadmap)

### Phase 1 – Core Service (MVP)

#### 🎯 Tujuan: Bisa demo basic key management + stack modern (Redis, RabbitMQ, gRPC).

- gRPC service dengan method:
  - GenerateKey (BTC dulu)
  - GetKey(id)
- Simpan hasil generate ke Redis (dengan TTL, misal 1 jam).
- Publish event ke RabbitMQ kalau ada key baru.
- Consumer RabbitMQ sederhana → log event ke console.
- Docker Compose untuk spin-up app + redis + rabbitmq.

👉 Di tahap ini, lo udah bisa show off gRPC, Redis, RabbitMQ, Docker.

### Phase 2 – Developer-Friendly

#### 🎯 Tujuan: Bikin lebih usable buat dev/crypto enthusiast.

- Tambah support Ethereum key (secp256k1).
- Derive BTC/ETH address dari public key.
- Generate QR code untuk public address.
- Tambah REST gateway (grpc-gateway) biar gampang diakses via Postman.
- CLI client sederhana (Go) buat call gRPC service.

👉 Lo udah bisa bilang: “Ini bukan cuma key generator, tapi bisa langsung dipakai buat dev testing wallet address”.

### Phase 3 – Security & Usability

#### 🎯 Tujuan: Mikirin security serius + UX.

- Support BIP39 mnemonic (seed phrase 12/24 kata).
- Support HD wallet derivation (BIP32/BIP44).
- Cold wallet mode → key langsung dihapus setelah dikirim (tidak tersimpan di Redis).
- Encrypt private key dengan password (AES).
- Export key ke JSON (mirip Ethereum keystore).

👉 Di sini keliatan lo ngerti crypto wallet standards + secure storage.

### Phase 4 – Advanced & Portfolio Killer

#### 🎯 Tujuan: Tambah nilai jual & bukti lo ngerti arsitektur besar.

- Fitur sign/verify message (praktis buat DApp dev).
- Webhook/Notification (misal integrasi ke Telegram/Discord via RabbitMQ consumer).
- Rate limiting (free tier vs premium).
- Audit log service → simpan semua aktivitas ke Postgres.
- JWT authentication untuk setiap API call.
- Deploy ke Kubernetes (opsional, biar makin next level).

👉 Di tahap ini, proyek lo udah bisa diposisikan sebagai mini SaaS crypto key management.

### Phase 5 – Show Off Mode 🚀

#### 🎯 Tujuan: Packaging buat portfolio & demo.

- Tulis dokumentasi + API examples (Swagger/OpenAPI untuk REST).
- Tambahin dashboard sederhana (Next.js/React) buat generate & lihat key (biar gak pure CLI).
- Publish di GitHub lengkap dengan docker-compose.yml & contoh client.
- (Optional) Deploy versi demo ke server/VPS gratis (misal fly.io / railway / render).

👉 Portfolio lo langsung keliatan solid: ada backend modern, event-driven, crypto-specific features, dokumentasi, bahkan demo online.

---

## 🛠 Tech Stack

- **Go** → main service
- **gRPC** → komunikasi antar service
- **Redis** → in-memory key storage (TTL)
- **RabbitMQ** → event bus untuk key events
- **Docker & Docker Compose** → containerization & orchestration

---

## 📐 Architecture (Phase 1)

```
+-------------+         +---------+          +-------------+
|   Client    | <-----> |  gRPC   | <------> |    Redis    |
| (CLI/REST)  |         | Service |          | (Key Cache) |
+-------------+         +---------+          +-------------+
        |                     |
        | (publish event)     |
        v                     v
   +-----------+        +-----------------+
   | RabbitMQ  | -----> | Audit Consumer  |
   | (exchange)|        | (log listener)  |
   +-----------+        +-----------------+
```

---

## ⚡ Getting Started

### 1. Clone repo

```bash
git clone https://github.com/yourusername/locksmith.git
cd locksmith
```

### 2. Jalankan Docker Compose

```bash
docker-compose up --build
```

### 3. Panggil gRPC Method

- `GenerateKey`
- `GetKey`

(Tutorial penggunaan via CLI/REST akan ditambahkan di fase berikutnya)

---

## 📅 Roadmap Demo

- Phase 1: Core MVP → Done ✅
- Phase 2: Multi-chain support + QR code → Next
- Phase 3: Security features (BIP39, HD Wallet) → Planned
- Phase 4: Advanced integration + SaaS tier → Planned
- Phase 5: Show off with frontend + demo deploy → Planned

---

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
