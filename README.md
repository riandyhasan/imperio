# 🏛️ Imperio

**Imperio** is a high-performance database simulation tool built in Go. It simulates **write**, **update**, and **delete** operations to a configured database, with customizable concurrency, rate limits, and schemas.

---

## ✨ Features

- 📦 Supports **PostgreSQL** and **MySQL** (extensible for more)
- ⚙️ YAML-based configuration
- 💥 Simulates operations at **N ops/sec**
- 🧵 Configurable **concurrency** and **duration**

---

## 📦 Installation

```bash
git clone https://github.com/riandyhasan/imperio.git
cd imperio
go build -o imperio ./cmd
```

---

## ⚙️ Configuration

Create a `config.yaml`:

```yaml
database: postgres
schema_file: ./schema.yaml
operations:
  - write
  - update
ops_per_second: 100
concurrency: 10
runner_duration: 30s
db_config:
  host: localhost
  port: '5432'
  user: imperio_user
  password: secret123
  dbname: imperio_db
  sslmode: disable
```

---

### 🧬 Schema File

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

```bash
./imperio -config=config.yaml
```

---

## 🧑‍💻 Author

Developed by [@riandyhasan](https://github.com/riandyhasan)
Open to contributions and feedback!

---
