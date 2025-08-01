# 🎬 Backend Test — Bioskop Ticketing API

### 👤 Fauzi Malik Nashrullah

A simple backend system for managing movie schedules in a national cinema chain. Built using **Go (Golang)**, **Gin**, and **PostgreSQL**, with basic JWT authentication.

---

## 📌 Fitur

- Login user
- CRUD data jadwal tayang bioskop
- Validasi relasi `film_id` dan `theater_id`
- JWT Authentication
- Dummy data otomatis saat start (admin + film + theater + schedule)

---

## ⚙️ Teknologi

| Teknologi  | Fungsi                   |
| ---------- | ------------------------ |
| Golang     | Bahasa pemrograman utama |
| Gin        | HTTP web framework       |
| GORM       | ORM untuk database       |
| PostgreSQL | Database utama           |
| JWT        | Otentikasi user          |
| Bcrypt     | Hash password            |

---

## 🗃️ Struktur Database

Tabel yang digunakan:

- `users`
- `films`
- `theaters`
- `schedules`

Model `Schedule` terhubung ke `Film` dan `Theater` via foreign key.  
Semua sudah di-`AutoMigrate`.

---

## 🚀 Cara Menjalankan

1. Buat database PostgreSQL, misalnya `cinema`
2. Edit konfigurasi koneksi di `config/db.go` → bagian `dsn`
3. Jalankan aplikasi:
   ```bash
   go run main.go
   ```

Saat pertama kali jalan, data dummy otomatis akan dibuat:

- User: `admin@example.com` / `password123`
- Film: `Avengers`
- Theater: `XXI Tunjungan`

---

## 🔐 Auth Endpoint

### POST `/login`

Login dan dapatkan JWT token.

Request:

```json
{
  "email": "admin@example.com",
  "password": "password123"
}
```

Response:

```json
{
  "token": "<JWT Token>"
}
```

```makefile
Authorization: Bearer <JWT Token>
```

---

## 📅 Schedule Endpoints (Protected)

### GET `/api/schedules`

Ambil semua jadwal tayang.

```json
[
  {
    "id": "...",
    "film_id": "...",
    "theater_id": "...",
    "show_time": "2025-08-02T19:30:00Z",
    "film": {
      "title": "Avengers"
    },
    "theater": {
      "name": "XXI Tunjungan"
    }
  }
]
```

---

### POST `/api/schedules`

Tambah jadwal tayang baru.

Request:

```json
{
  "film_id": "<UUID>",
  "theater_id": "<UUID>",
  "show_time": "2025-08-02T19:30:00Z"
}
```

Response:

```json
{
  "id": "...",
  "film": { ... },
  "theater": { ... }
}
```

---

### PUT `/api/schedules/:id`

Update jadwal.

---

### DELETE `/api/schedules/:id`

Hapus jadwal.

Response:

```json
{
  "message": "Deleted successfully"
}
```

---

## 🧪 Dummy Akun & Data

| Tipe     | Data                    |
| -------- | ----------------------- |
| Email    | `admin@example.com`     |
| Password | `password123`           |
| Film     | Avengers                |
| Theater  | XXI Tunjungan, Surabaya |

---

## 🙋 Author

Fauzi Malik Nashrullah
