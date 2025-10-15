# 🚀 Setup Project

## 1. Download & Install Go

**Windows/Mac:**
- Download dari: https://go.dev/dl/
- Install seperti biasa
- Verifikasi: `go version`

**Linux:**
```bash
sudo apt install golang-go
```

---

## 2. Clone Project

```bash
git clone https://github.com/aryarobyy/bank-soal
cd bank-soal/backend
```

---

## 3. Install Dependencies

```bash
go mod download
```

atau

```bash
go mod tidy
```

---

## 4. Setup Environment

### Buat file `.env`

```env
SERVER_PORT=8080
JWT_SECRET=
JWT_EXPIRED=

DB_USERNAME=root
DB_PASSWORD=
DB_NAME=
DB_HOST=
DB_PORT=

DB_MAX_IDLE_CONNS=
DB_MAX_OPEN_CONNS=
DB_CONN_MAX_LIFETIME=
```

---

## 5. Run Application

```bash
go run main.go
```

---

## 🎯 Done!

Server berjalan di: `http://localhost:8080`

---

## 📝 Common Commands

```bash
go mod tidy          # Install/update dependencies
go run main.go       # Run aplikasi
```