# 🏛️ Imperio

**Imperio** is a high-performance, configurable **database operation simulator** built in Go.  
It simulates concurrent **write**, **update**, and **delete** operations on your database at controlled rates, helping you **stress-test**, **benchmark**, or **emulate real-world workloads**.

---

## ✨ Features

- 📦 Supports **PostgreSQL** and **MySQL** (extensible with Strategy Pattern)
- 🧵 Configurable **concurrency** and **operations per second**
- ⚙️ YAML-based config and schema file
- ⏱️ Supports fixed-duration or infinite simulation

---

## 📦 Installation

```bash
git clone https://github.com/riandyhasan/imperio.git
cd imperio
go build -o imperio ./cmd
```

---

## ⚙️ Configuration

Create a `config.yaml` file in your project root:

```yaml
database: postgres
schema_file: ./schema.yaml
operations:
  - write
  - update
ops_per_second: 100
concurrency: 10
runner_duration: 30s # 0 or negative for infinite run
db_config:
  host: localhost
  port: '5432'
  user: imperio_user
  password: secret123
  dbname: imperio_db
  sslmode: disable
```

---

### 🧬 Schema File (`schema.yaml`)

```yaml
table: users
fields:
  id: int
  name: string
  email: string
  created_at: timestamp
```

---

## 🚀 Running Imperio

### Run Locally

```bash
./imperio --config=config.yaml
```

### Run with Docker

```bash
docker-compose up --build
```

> Use `config.postgres.yaml` or `config.mysql.yaml` to switch databases.

---

## 🧪 Development

### Format & Lint

```bash
make fmt
make lint
```

### Run Unit Tests

```bash
make test
```

### Build & Run Binary

```bash
make build
make run
```

---

## 🐳 Docker

### Build Image

```bash
make docker
```

### Run Compose

```bash
make docker-run
```

---

## 👤 Author

Developed by [@riandyhasan](https://github.com/riandyhasan)  
🛠️ Open to contributions, PRs, or feedback!
