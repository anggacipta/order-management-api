# Order Management API

Order Management API adalah aplikasi backend sederhana untuk manajemen produk, pemesanan, dan autentikasi user berbasis Golang (Gin, GORM, SQLite).

## Fitur
- Register & Login (JWT Auth)
- CRUD Produk (admin only)
- Pemesanan produk oleh customer (stok otomatis berkurang)
- Riwayat pesanan customer
- Validasi & error handling
- Dokumentasi API (Swagger/OpenAPI)
- Unit testing

## Instalasi & Menjalankan Project

### 1. Clone repository
```bash
git clone https://github.com/anggacipta/order-management-api.git
cd order-management-api
```

### 2. Install dependency
```bash
go mod tidy
```

### 3. Jalankan aplikasi
```bash
go run main.go
```

Aplikasi akan berjalan di `http://localhost:8080`

### 4. Testing
```bash
go test ./tests
```

### 5. Dokumentasi API
- File dokumentasi Swagger ada di `swagger.yaml`.
- Bisa dibuka di [Swagger Editor](https://editor.swagger.io/).

## Asumsi
- User dengan role `admin` hanya bisa dibuat lewat endpoint `/register-admin`.
- Semua endpoint `/admin/*` hanya bisa diakses admin (dengan JWT).
- Semua endpoint `/orders` hanya bisa diakses user login (JWT).
- Database default: SQLite (file `order.db`).

## Struktur Project
- `main.go` : Entry point aplikasi
- `models/` : Model database & koneksi
- `controllers/` : Handler endpoint
- `dto/` : Data Transfer Object (request struct)
- `middlewares/` : Middleware JWT & role
- `routes/` : Routing aplikasi
- `helpers/` : Helper error handling
- `tests/` : Unit test
- `swagger.yaml` : Dokumentasi API

---

Silakan edit README ini sesuai kebutuhan dan tambahkan instruksi lain jika ada perubahan pada project.
